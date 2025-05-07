package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
)

func HandleError(w http.ResponseWriter, err error) {
	log.Printf("Error: %v", err)

	w.Header().Set("Content-Type", "application/json")

	if appError, ok := err.(*errors.AppError); ok {
		w.WriteHeader(appError.GetStatusCode())
		json.NewEncoder(w).Encode(appError.ToHTTPResponse())
		return
	}

	// Erro genérico (não tratado)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]string{
			"code":    "internal_error",
			"message": "An unexpected error occurred",
		},
	})
}
