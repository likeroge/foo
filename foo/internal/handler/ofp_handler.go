package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ego.dev21/greetings/internal/entities"
	"ego.dev21/greetings/internal/repository"
	"ego.dev21/greetings/internal/usecases"
	"ego.dev21/greetings/internal/utils"
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
			log.Println("GetOFPInfoById error during Atoi: ", err)
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
			log.Println("DeleteOFPInfoById error during Atoi: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		h.Repositories.OFPRpository.DeleteOFPInfoById(intVal)
	}
}

func (h *OfpHandler) GetAllOFPInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		ofpInfo, err := h.Repositories.OFPRpository.GetAllOFPInfo()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(ofpInfo)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *OfpHandler) PostOfpToBackend(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Println("PostOfpToBackend error during FormFile: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer file.Close()

		pdfContent, err := utils.ReadPDFContent(file, fileHeader)
		if err != nil {
			log.Println("PostOfpToBackend -> ReadPDFContent error: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ofpParser := usecases.NewOFPParser(pdfContent.PdfContent)
		parsedOfp, err := ofpParser.ParseOfp()
		if err != nil {
			log.Println("PostOfpToBackend -> ParseOfp error: ", err)
			w.Header().Set("Content-Type", "application/json")
			e := entities.NewAPIError(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
			return
		}
		_, err = h.Repositories.OFPRpository.CreateOFPInfo(parsedOfp)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			e := entities.NewAPIError(err.Error())
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(e)
			return
		}

		pdfContent.ParsedOfp = parsedOfp

		byteContent, err := json.Marshal(parsedOfp)
		if err != nil {
			log.Println("PostOfpToBackend -> Marshal(parsedOfp) error: ", err)
			w.Write([]byte(err.Error()))
			return
		}
		fp := usecases.NewFileProducer(&byteContent)
		fp.FileName = "ofp_data " + parsedOfp.RegNumber + " " + parsedOfp.FlightNumber + " " + parsedOfp.IcaoFrom + " " + parsedOfp.IcaoTo
		fp.FileType = "txt"
		fp.SendFileViaHttp(w)
	}
}
