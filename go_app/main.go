package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// GyroscopeData represents the expected gyroscope payload
type GyroscopeData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// GPSData represents the expected GPS payload
type GPSData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

// PhotoData represents the expected photo payload
type PhotoData struct {
	Filename string `json:"filename"`
	Data     string `json:"data"` // base64 encoded string or any string data
}

func main() {
	http.HandleFunc("/telemetry/gyroscope", gyroscopeHandler)
	http.HandleFunc("/telemetry/gps", gpsHandler)
	http.HandleFunc("/telemetry/photo", photoHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func gyroscopeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data GyroscopeData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	// Validate required fields (all must be present and non-zero)
	if data.X == 0 && data.Y == 0 && data.Z == 0 {
		http.Error(w, "Missing or invalid gyroscope data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Gyroscope data received successfully"}`))
}

func gpsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data GPSData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	// Validate required fields (latitude and longitude must be non-zero)
	if data.Latitude == 0 || data.Longitude == 0 {
		http.Error(w, "Missing or invalid GPS data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"GPS data received successfully"}`))
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data PhotoData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate required fields (filename and data must be non-empty)
	if data.Filename == "" || data.Data == "" {
		http.Error(w, "Missing photo data fields", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Photo data received successfully"}`))
}
