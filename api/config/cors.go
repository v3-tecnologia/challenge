package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         300,
	}))
}
