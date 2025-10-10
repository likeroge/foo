package usecases

import (
	"fmt"
	"net/http"
)

type FileProducer struct {
	FileType      string
	FileName      string
	BinaryContent *[]byte
}

func NewFileProducer(binaryContent *[]byte) *FileProducer {
	return &FileProducer{FileType: "txt", FileName: "fileName", BinaryContent: binaryContent}
}

func (f *FileProducer) SendFileViaHttp(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Disposition", "attachment; filename="+f.FileName+"."+f.FileType)

	_, err := w.Write(*f.BinaryContent)
	if err != nil {
		fmt.Println(err)
	}
}

// func (f *FileProducer) CreateTXTFileInMemory() (*os.File, error) {
// 	file, err := os.Create(f.FileName + "." + f.FileType)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	file.Write(*f.BinaryContent)
// 	return file, nil
// }
