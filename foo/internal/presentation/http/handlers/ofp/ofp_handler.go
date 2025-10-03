package ofp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ego.dev21/greetings/internal/entities"
	OfpParser "ego.dev21/greetings/internal/usecases/ofp"
	"ego.dev21/greetings/internal/utils"
)

func PostOfpToBackend(w http.ResponseWriter, r *http.Request) {
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
		ofpParser := OfpParser.NewOFPParser(pdfContent)
		parsedOfp, err := ofpParser.ParseOfp()
		if err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			e := entities.NewAPIError(err.Error())
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(e)
			return
		}
		fmt.Println(parsedOfp)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(parsedOfp)
	}
}
