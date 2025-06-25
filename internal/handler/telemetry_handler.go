package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/core/usecase"
	"github.com/yanvic/challenge/infra/database/dynamo"
	"github.com/yanvic/challenge/internal/queue"
)

func HandlerGyroscope(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data entity.Gyroscope
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := usecase.ValidateGyroscope(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := dynamo.SaveGyro(*data.X, *data.Y, *data.Z, data.DeviceID, data.Timestamp, data)
	if err != nil {
		log.Printf("Erro ao salvar no DynamoDB: %v\n", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Gyroscope data saved"))
}

func HandlerGPS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data entity.GPS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := usecase.ValidateGPS(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := dynamo.SaveGps(*data.Latitude, *data.Longitude, data.DeviceID, data.Timestamp, data)
	if err != nil {
		log.Printf("Erro ao salvar no DynamoDB: %v\n", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GPS data received"))
}

var PublishImageFunc = queue.PublishImage

func HandlerPhoto(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error processing image", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image not uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	image, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading image", http.StatusInternalServerError)
		return
	}

	deviceID := r.FormValue("device_id")
	timestamp := r.FormValue("timestamp")
	if timestamp == "" {
		timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	payload := entity.Photo{
		Image:     image,
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Error serializing payload", http.StatusInternalServerError)
		return
	}

	err = PublishImageFunc(jsonPayload)
	if err != nil {
		http.Error(w, "Error publishing to queue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image uploaded successfully"))
}
