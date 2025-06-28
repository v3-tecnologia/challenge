package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"v3-test/internal/dtos/telemetriesDtos"

	"github.com/stretchr/testify/assert"
)

func TestCreateGpsIntegration(t *testing.T) {
	payload := telemetriesDtos.CreateGpsDto{
		Latitude:  F64(12.34),
		Longitude: F64(56.78),
	}

	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestInvalidCreateGpsIntegration(t *testing.T) {
	payload := telemetriesDtos.CreateGpsDto{
		Latitude:  nil,
		Longitude: F64(56.78),
	}

	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
