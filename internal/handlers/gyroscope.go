package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/bielgennaro/v3-challenge-cloud/internal/repository"
)

type GyroscopeRequest struct {
	MacAddress string    `json:"mac_address"`
	AxisX      float64   `json:"axis_x"`
	AxisY      float64   `json:"axis_y"`
	AxisZ      float64   `json:"axis_z"`
	Timestamp  time.Time `json:"timestamp"`
}

func HandleGyroscopeData(w http.ResponseWriter, r *http.Request) {
	var req GyroscopeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, errors.NewErrorBadRequest("invalid_json", "Failed to parse request body"))
		return
	}

	gyroscope, err := model.NewGyroscopeBuilder().
		WithMacAddress(req.MacAddress).
		WithAxisValues(req.AxisX, req.AxisY, req.AxisZ).
		WithTimestamp(req.Timestamp).
		Build()

	if err != nil {
		HandleError(w, err)
		return
	}

	if err := repository.SaveGyroscope(gyroscope); err != nil {
		HandleError(w, errors.NewErrorInternal("database_error", "Failed to save gyroscope data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Gyroscope data saved successfully",
		"id":      gyroscope.ID,
	})
}
