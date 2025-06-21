package main

import (
	"log"
	"net/http"

	"github.com/yanvic/challenge/internal/handler"
)

func main() {
	http.HandleFunc("/telemetry/gyroscope", handler.HandlerGyroscope)
	http.HandleFunc("/telemetry/gps", handler.HandlerGPS)
	http.HandleFunc("/telemetry/photo", handler.HandlerPhoto)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
