package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockAddGpsSuccess(gps models.GpsModel) error {
	return nil
}

func mockAddGpsError(gps models.GpsModel) error {
	return fmt.Errorf("database error")
}

func TestSaveGps_Happy(t *testing.T) {
	routes.SetGpsService(mockAddGpsSuccess)

	r := gin.Default()
	r.POST("/telemetry/gps", routes.SaveGps)

	gps := models.GpsModel{
		Latitude:  40.7128,
		Longitude: -74.0060,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	jsonGps, _ := json.Marshal(gps)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(jsonGps))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, but got %v", http.StatusCreated, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["status"] != "GPS Saved Successfully" {
		t.Errorf("Expected success message 'GPS Saved Successfully', but got %v", response["status"])
	}
}

func TestSaveGps_BadRequest(t *testing.T) {
	routes.SetGpsService(mockAddGpsSuccess)

	r := gin.Default()
	r.POST("/telemetry/gps", routes.SaveGps)

	gps := map[string]interface{}{"bad": "bad"}

	jsonGps, _ := json.Marshal(gps)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(jsonGps))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, but got %v", http.StatusBadRequest, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["error"] != "Incorrect or missing parameters" {
		t.Errorf("Expected success message 'Incorrect or missing parameters', but got %v", response["error"])
	}
}

func TestSaveGps_InternalServerError(t *testing.T) {
	routes.SetGpsService(mockAddGpsError)

	r := gin.Default()
	r.POST("/telemetry/gps", routes.SaveGps)

	gps := models.GpsModel{
		Latitude:  40.7128,
		Longitude: -74.0060,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	jsonGps, _ := json.Marshal(gps)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(jsonGps))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["error"] != "Error saving gps to database: database error" {
		t.Errorf("Expected error message 'Error saving gps to database: database error', but got %v", response["error"])
	}
}
