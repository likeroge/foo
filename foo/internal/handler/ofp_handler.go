package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ego.dev21/greetings/internal/entities"
	"ego.dev21/greetings/internal/repository"
	"ego.dev21/greetings/internal/utils"

	OfpParser "ego.dev21/greetings/internal/usecases/ofp_use_cases"
)

type OfpHandler struct {
	Repositories *repository.Repositories
}

func NewOfpHandler(repositories *repository.Repositories) *OfpHandler {
	return &OfpHandler{
		Repositories: repositories,
	}
}

func (h *OfpHandler) GetOFPInfoById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		pathVal := r.PathValue("id")
		intVal, err := strconv.Atoi(pathVal)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		ofpInfo := h.Repositories.OFPRpository.GetOFPInfoById(intVal)
		json.NewEncoder(w).Encode(ofpInfo)
	}
}

func (h *OfpHandler) DeleteOFPInfoById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		pathVal := r.PathValue("id")
		intVal, err := strconv.Atoi(pathVal)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		h.Repositories.OFPRpository.DeleteOFPInfoById(intVal)
	}
}

func (h *OfpHandler) GetAllOFPInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		ofpInfo, err := h.Repositories.OFPRpository.GetAllOFPInfo()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(ofpInfo)
	}
}

func (h *OfpHandler) PostOfpToBackend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		pdfContent, err := utils.ReadPDFContent(file, fileHeader)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		ofpParser := OfpParser.NewOFPParser(pdfContent.PdfContent)
		parsedOfp, err := ofpParser.ParseOfp()
		if err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			e := entities.NewAPIError(err.Error())
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(e)
			return
		}
		lastInsertedId := h.Repositories.OFPRpository.CreateOFPInfo(parsedOfp)
		fmt.Println(lastInsertedId)

		pdfContent.ParsedOfp = parsedOfp
		fmt.Println(parsedOfp)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parsedOfp)
	}
}
