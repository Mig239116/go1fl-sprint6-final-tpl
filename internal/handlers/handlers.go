package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const IndexPath = "../index.html"

func IndexHanlder(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusMethodNotAllowed)
		return
	}
	htmlContent, err := os.ReadFile(IndexPath)
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlContent)
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusMethodNotAllowed)
		return
	}
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "Cant get the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := os.ReadFile("../" + header.Filename)
	if err != nil {
		http.Error(w, "Failed to read the file", http.StatusBadRequest)
		return
	}
	text := string(data)
	newText, err := service.ConvertMessage(text)
	if err != nil {
		http.Error(w, "Failed to convert data", http.StatusInternalServerError)
		return
	}
	name := fmt.Sprintf("../file_%s%s", time.Now().UTC().Format("20060102_150405"), filepath.Ext("*.txt"))
	newFile, err := os.Create(name)
	if err != nil {
		http.Error(w, name, http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	_, err = newFile.WriteString(newText)
	if err != nil {
		http.Error(w, "Failed to write result to file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, newText)
}
