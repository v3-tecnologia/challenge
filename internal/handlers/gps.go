package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/bielgennaro/v3-challenge-cloud/internal/repository"
)

type GPSRequest struct {
	MacAddress string    `json:"mac_address"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Timestamp  time.Time `json:"timestamp"`
}

func HandleGPSData(w http.ResponseWriter, r *http.Request) {
	var req GPSRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, errors.NewErrorBadRequest("invalid_json", "Failed to parse request body"))
		return
	}

	lat := req.Latitude
	lng := req.Longitude

	gps, err := model.NewGPSBuilder().
		WithMacAddress(req.MacAddress).
		WithCoordinates(&lat, &lng).
		WithTimestamp(req.Timestamp).
		Build()

	if err != nil {
		log.Printf("Error building GPS model: %v", err)
		HandleError(w, err)
		return
	}

	if err := repository.SaveGPS(gps); err != nil {
		log.Printf("Error saving GPS: %v", err)
		HandleError(w, errors.NewErrorInternal("database_error", "Failed to save GPS data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "GPS data saved successfully",
		"id":      gps.ID,
	})
}
