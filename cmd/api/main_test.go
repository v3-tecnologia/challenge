package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bielgennaro/v3-challenge-cloud/internal/routes"
	"github.com/gorilla/mux"
)

func TestTelemetryRoutes(t *testing.T) {
	router := mux.NewRouter()
	routes.SetupTelemetryRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("esperado status %d, recebido %d", http.StatusOK, status)
	}

	expected := "ok"
	if rr.Body.String() != expected {
		t.Errorf("esperado body '%s', recebido '%s'", expected, rr.Body.String())
	}
}
