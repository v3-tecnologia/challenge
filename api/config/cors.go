package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupCORS(router *chi.Mux) {
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         300,
	}

	router.Use(cors.Handler(corsOptions))
}
