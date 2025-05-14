package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/martinsrenan/challenge/internal/model"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data model.GyroscopeData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if data.X == 0 && data.Y == 0 && data.Z == 0 {
		http.Error(w, "Missing or invalid gyroscope data", http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("INSERT INTO gyroscope (x, y, z) VALUES ($1, $2, $3)", data.X, data.Y, data.Z)
	if err != nil {
		http.Error(w, "Failed to save gyroscope data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gyroscope data received successfully"}`))
}

func (h *Handler) GPSHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data model.GPSData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("INSERT INTO gps (latitude, longitude, altitude) VALUES ($1, $2, $3)", data.Latitude, data.Longitude, data.Altitude)
	if err != nil {
		http.Error(w, "Failed to save GPS data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"GPS data received successfully"}`))
}

func (h *Handler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data model.PhotoData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if data.Filename == "" || data.Data == "" {
		http.Error(w, "Missing or invalid photo data", http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("INSERT INTO photos (filename, data) VALUES ($1, $2)", data.Filename, data.Data)
	if err != nil {
		http.Error(w, "Failed to save photo data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Photo data received successfully"}`))
}
