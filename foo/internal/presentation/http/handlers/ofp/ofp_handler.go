package ofp

import (
	"fmt"
	"log"
	"net/http"

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
			log.Fatal(err)
		}
		ofpParser := OfpParser.NewOFPParser(pdfContent)
		ofpParser.ParseOfp()

	}
}
