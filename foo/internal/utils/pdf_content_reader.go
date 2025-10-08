package utils

import (
	"fmt"
	"mime/multipart"

	"ego.dev21/greetings/internal/entities"
	"github.com/ledongthuc/pdf"
)

type PdfContentReader struct {
	PdfContent string
	Error      error
	ParsedOfp  *entities.OFP
}

// func ReadPDFContent(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
func ReadPDFContent(file multipart.File, fileHeader *multipart.FileHeader) (*PdfContentReader, error) {

	buf := make([]byte, fileHeader.Size)

	fmt.Println(fileHeader.Filename)
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)

	pdfReader, err := pdf.NewReader(file, fileHeader.Size)
	if err != nil {
		// return "", err
		return nil, err

	}
	reader, err := pdfReader.GetPlainText()
	if err != nil {
		return nil, err
		// return "", err
	}
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
		// return "", err
	}
	return &PdfContentReader{PdfContent: string(buf), Error: nil, ParsedOfp: nil}, nil
	// return string(buf), err
}
