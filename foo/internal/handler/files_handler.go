package handler

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"ego.dev21/greetings/internal/repository"
)

type FilesHandler struct {
	Repositories *repository.Repositories
}

type PageData struct {
	Title     string
	Items     []string
	TimeStamp string
}

func NewFilesHandler(repositories *repository.Repositories) *FilesHandler {
	return &FilesHandler{
		Repositories: repositories,
	}
}

func (h *FilesHandler) GetHelloTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	pageData := &PageData{
		Title:     "This is Template Title",
		Items:     []string{"Item 1", "Item 2", "Item 3"},
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	t, err := template.ParseFiles("templates/template1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, pageData)
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
