package main

import (
	"log"
	"net/http"

	"github.com/martinsrenan/challenge/internal/config"
	"github.com/martinsrenan/challenge/internal/db"
	"github.com/martinsrenan/challenge/internal/handler"
)

func main() {
	cfg := config.Load()
	dbConn := db.Connect(cfg.DatabaseURL)
	defer dbConn.Close()

	h := handler.New(dbConn)

	http.HandleFunc("/telemetry/gyroscope", h.GyroscopeHandler)
	http.HandleFunc("/telemetry/gps", h.GPSHandler)
	http.HandleFunc("/telemetry/photo", h.PhotoHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
