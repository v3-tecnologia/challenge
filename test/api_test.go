package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"v3-backend-challenge/src/dto"
)

func TestHandleGyroscope(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	payload := dto.Gyroscope{
		AxisX: 1.2,
		AxisY: -0.8,
		AxisZ: 0.4,
		BaseTelemetry: dto.BaseTelemetry{
			DateTimeCollected: time.Now(),
			MacAddr:           "00:11:22:33:44:55",
		},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandleGPS(t *testing.T) {
	db := SetupTestDB()
	router := SetupRouter(db)

	payload := dto.GPS{
		BaseTelemetry: dto.BaseTelemetry{
			DateTimeCollected: time.Now(),
			MacAddr:           "00:11:22:33:44:55",
		},
		Latitude:  30,
		Longitude: 4,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandlePhoto(t *testing.T) {

}
