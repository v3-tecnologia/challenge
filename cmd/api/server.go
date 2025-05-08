package main

import (
	"log"
	"net/http"
	"time"

	env "github.com/bielgennaro/v3-challenge-cloud/config"
	"github.com/bielgennaro/v3-challenge-cloud/config/middleware"
	"github.com/bielgennaro/v3-challenge-cloud/internal/routes"
	"github.com/gorilla/mux"
)

func RunServer() *http.Server {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	routes.SetupTelemetryRoutes(router)

	port := env.GetEnv("PORT", "8080")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Server listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	return srv
}
