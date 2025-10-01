package utils

import (
	"fmt"
	"mime/multipart"

	"github.com/ledongthuc/pdf"
)

func ReadPDFContent(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	buf := make([]byte, fileHeader.Size)

	fmt.Println(fileHeader.Filename)
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)

	pdfReader, err := pdf.NewReader(file, fileHeader.Size)
	if err != nil {
		return "", err
	}
	reader, err := pdfReader.GetPlainText()
	if err != nil {
		return "", err
	}
	_, err = reader.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf), err
}
