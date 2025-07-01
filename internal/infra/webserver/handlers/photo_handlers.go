package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/igorlopes88/desafio-v3/internal/dto"
	"github.com/igorlopes88/desafio-v3/internal/entity"
	"github.com/igorlopes88/desafio-v3/internal/infra/database"
)

type PhotoHandler struct {
	PhotoDB database.PhotoInterface
}

func NewPhotoHandler(db database.PhotoInterface) *PhotoHandler {
	return &PhotoHandler{
		PhotoDB: db,
	}
}

func (p *PhotoHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2 << 20)
	if err != nil {
		http.Error(w, "limit size image is 2 mb", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	if fileHeader.Size > (1024 * 1024 * 2) {
		http.Error(w, "limit size image is 2 mb", http.StatusBadRequest)
		return
	}

	macAddress := r.FormValue("mac_address")
	if macAddress == "" {
		http.Error(w, "mac_address is required", http.StatusBadRequest)
		return
	}

	// VERIFICATION TYPE FILE
	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".gif" && ext != ".png" {
		http.Error(w, "file type not accepted", http.StatusBadRequest)
		return
	}

	// RENAME FILE
	timestamp := time.Now().Format("20060102-150405")
	fileName := fmt.Sprintf("%s-%s%s", timestamp, uuid.New().String(), ext)
	filePath := filepath.Join("./uploads/", fileName)

	// SAVE FILE
	f, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()

	path, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer path.Close()

	_, err = io.Copy(path, f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := dto.RegisterPhotoInput{
		UserId:     uuid.New(),
		MacAddress: macAddress,
		Image:      filePath,
		TimeStamp:  "2025-06-30 10:00:10",
	}

	g, err := entity.NewPhoto(input.UserId, input.MacAddress, input.Image, input.TimeStamp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = p.PhotoDB.Register(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
