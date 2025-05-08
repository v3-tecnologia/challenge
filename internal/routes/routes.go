package routes

import (
	"github.com/gorilla/mux"

	handlers "github.com/bielgennaro/v3-challenge-cloud/internal/handlers"
)

func SetupTelemetryRoutes(router *mux.Router) {
	telemetryRouter := router.PathPrefix("/telemetry").Subrouter()

	telemetryRouter.HandleFunc("/gyroscope", handlers.HandleGyroscopeData).Methods("POST")
	telemetryRouter.HandleFunc("/gps", handlers.HandleGPSData).Methods("POST")
	telemetryRouter.HandleFunc("/photo", handlers.HandlePhotoData).Methods("POST")

	router.HandleFunc("/health", handlers.HandleHealthCheck).Methods("GET")
}
