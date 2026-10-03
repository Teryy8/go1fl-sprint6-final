package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(res, "Read form error", http.StatusInternalServerError)
		return
	}
	defer req.MultipartForm.RemoveAll()

	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "File recive error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "File read error", http.StatusInternalServerError)
		return
	}

	result, err := service.Translater(string(data))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().String() +
		filepath.Ext(header.Filename)

	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		http.Error(res, "File save error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(res, result)
}
