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

func TestCreateGyroscopeIntegration(t *testing.T) {
	payload := telemetriesDtos.CreateGyroscopeDto{
		X: F64(1.23),
		Y: F64(4.56),
		Z: F64(7.89),
	}

	jsonPayload, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestInvalidCreateGyroscopeIntegration(t *testing.T) {
	payload := telemetriesDtos.CreateGyroscopeDto{
		X: nil,
		Y: F64(4.56),
		Z: F64(7.89),
	}

	jsonPayload, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
