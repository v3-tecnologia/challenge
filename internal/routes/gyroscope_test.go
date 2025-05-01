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

func mockAddGyroscopeSuccess(gyroscope models.GyroscopeModel) error {
	return nil
}

func mockAddGyroscopeError(gyroscope models.GyroscopeModel) error {
	return fmt.Errorf("database error")
}

func TestSaveGyroscope_Happy(t *testing.T) {
	routes.SetGyroscopeService(mockAddGyroscopeSuccess)

	r := gin.Default()
	r.POST("/telemetry/gyroscope", routes.SaveGyroscope)

	gyroscope := models.GyroscopeModel{
		X:         40.7128,
		Y:         40.7128,
		Z:         40.7128,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	jsonGyroscope, _ := json.Marshal(gyroscope)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(jsonGyroscope))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, but got %v", http.StatusCreated, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["status"] != "Gyroscope Saved Successfully" {
		t.Errorf("Expected success message 'Gyroscope Saved Successfully', but got %v", response["status"])
	}
}

func TestSaveGyroscope_BadRequest(t *testing.T) {
	routes.SetGyroscopeService(mockAddGyroscopeSuccess)

	r := gin.Default()
	r.POST("/telemetry/gyroscope", routes.SaveGyroscope)

	gyroscope := map[string]interface{}{"bad": "bad"}

	jsonGyroscope, _ := json.Marshal(gyroscope)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(jsonGyroscope))
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

func TestSaveGyroscope_InternalServerError(t *testing.T) {
	routes.SetGyroscopeService(mockAddGyroscopeError)

	r := gin.Default()
	r.POST("/telemetry/gyroscope", routes.SaveGyroscope)

	gyroscope := models.GyroscopeModel{
		X:         40.7128,
		Y:         40.7128,
		Z:         40.7128,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	jsonGyroscope, _ := json.Marshal(gyroscope)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(jsonGyroscope))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["error"] != "Error saving gyroscope to database: database error" {
		t.Errorf("Expected error message 'Error saving gyroscope to database: database error', but got %v", response["error"])
	}
}
