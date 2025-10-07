package handler

import (
	"net/http"
	"os"

	"ego.dev21/greetings/internal/repository"
)

type FilesHandler struct {
	Repositories *repository.Repositories
}

func NewFilesHandler(repositories *repository.Repositories) *FilesHandler {
	return &FilesHandler{
		Repositories: repositories,
	}
}

func (h *FilesHandler) SendFile(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	// w.Header().Set("Content-Type", "text/plain")
	//read file

	fileContent, err := os.ReadFile("go.mod")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := "imp.txt"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	// bytes.NewReader(fileContent)
	w.Write([]byte("File content: " + string(fileContent)))

	// w.Write([]byte("Hello, World!"))

}
