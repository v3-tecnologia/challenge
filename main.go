package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// GyroscopeData represents the expected gyroscope payload
type GyroscopeData struct {
	ID int     `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
}

// GPSData represents the expected GPS payload
type GPSData struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

// PhotoData represents the expected photo payload
type PhotoData struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Data     string `json:"data"` // base64 encoded string or any string data
}

var db *sql.DB

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

	// Validate required fields
	if data.Filename == "" || data.Data == "" {
		http.Error(w, "Missing or invalid photo data", http.StatusBadRequest)
		return
	}

	// Save to database
	_, err := db.Exec("INSERT INTO photos (filename, data) VALUES ($1, $2)", data.Filename, data.Data)
	if err != nil {
		http.Error(w, "Failed to save photo data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Photo data received successfully"}`))
}

func main() {
	var err error
	// Connect to the PostgreSQL database
	connStr := "user=yourusername dbname=yourdbname sslmode=disable" // Update with your credentials
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables if they do not exist
	createTables()

	http.HandleFunc("/telemetry/gyroscope", gyroscopeHandler)
	http.HandleFunc("/telemetry/gps", gpsHandler)
	http.HandleFunc("/telemetry/photo", photoHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTables() {
	// Create table for gyroscope data
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS gyroscope (
        id SERIAL PRIMARY KEY,
        x FLOAT NOT NULL,
        y FLOAT NOT NULL,
        z FLOAT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create table for GPS data
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gps (
        id SERIAL PRIMARY KEY,
        latitude FLOAT NOT NULL,
        longitude FLOAT NOT NULL,
        altitude FLOAT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create table for photo data
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS photos (
        id SERIAL PRIMARY KEY,
        filename TEXT NOT NULL,
        data TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}
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
	// Validate required fields
	if data.X == 0 && data.Y == 0 && data.Z == 0 {
		http.Error(w, "Missing or invalid gyroscope data", http.StatusBadRequest)
		return
	}

	// Save to database
	_, err := db.Exec("INSERT INTO gyroscope (x, y, z) VALUES ($1, $2, $3)", data.X, data.Y, data.Z)
	if err != nil {
		http.Error(w, "Failed to save gyroscope data", http.StatusInternalServerError)
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
}
