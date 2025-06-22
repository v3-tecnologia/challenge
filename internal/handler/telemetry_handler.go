package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/core/usecase"
	"github.com/yanvic/challenge/infra/database/dynamo"
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

func HandlerPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data entity.Photo
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := usecase.ValidatePhoto(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := dynamo.SavePhoto(data.ImageBase64, data.DeviceID, data.Timestamp, data)
	if err != nil {
		log.Printf("Erro ao salvar no DynamoDB: %v\n", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Photo data received"))
}
