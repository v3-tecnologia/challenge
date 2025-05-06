package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	route "github.com/mkafonso/go-cloud-challenge/api/routes"
	recognition "github.com/mkafonso/go-cloud-challenge/recognition/provider"
	storage "github.com/mkafonso/go-cloud-challenge/storage/provider"

	"github.com/mkafonso/go-cloud-challenge/api/config"
)

func main() {
	log.Printf("🚀 Starting API on port %s\n", "8080")

	router := chi.NewRouter()

	// Setup CORS
	config.SetupCORS(router)

	// Setup Postgres
	pool := config.SetupPostgres()
	repositories := config.SetupRepositories(pool)

	// Mocks para Rekognition e Upload
	recognizer := recognition.NewInMemoryFaceRecognition()
	storage := storage.NewInMemoryPhotoStorage()

	// Mount telemetry routes
	router.Mount("/telemetry", route.TelemetryModuleRouter(
		repositories.GPSRepo,
		repositories.GyroscopeRepo,
		repositories.PhotoRepo,
		storage,
		recognizer,
	))

	startRestHTTPServer(router)
}

func startRestHTTPServer(router http.Handler) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
