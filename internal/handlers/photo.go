package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/bielgennaro/v3-challenge-cloud/internal/repository"
)

type PhotoRequest struct {
	MacAddress string    `json:"mac_address"`
	FileUrl    string    `json:file_url`
	Timestamp  time.Time `json:"timestamp"`
}

func HandlePhotoData(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		HandleError(w, errors.NewErrorBadRequest("invalid_form", "Request must be multipart/form-data"))
		return
	}

	macAddress := r.FormValue("mac_address")
	timestampStr := r.FormValue("timestamp")
	fileUrl := r.FormValue("file_url")
	if macAddress == "" || timestampStr == "" || fileUrl == "" {
		HandleError(w, errors.NewErrorBadRequest("missing_fields", "mac_address, timestamp and file_url are required"))
		return
	}

	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		HandleError(w, errors.NewErrorBadRequest("invalid_timestamp", "Timestamp must be in RFC3339 format"))
		return
	}

	photo, err := model.NewPhotoBuilder().
		WithMacAddress(macAddress).
		WithRecognitionStatus(false).
		WithFileUrl(fileUrl).
		WithTimestamp(timestamp).
		Build()

	if err != nil {
		HandleError(w, err)
		return
	}

	if err := repository.SavePhoto(photo); err != nil {
		HandleError(w, errors.NewErrorInternal("database_error", "Failed to save photo data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Photo uploaded successfully",
		"id":      photo.ID,
	})
}
