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

func mockAddPhotoSuccess(photo models.PhotoModel) error {
	return nil
}

func mockAddPhotoError(photo models.PhotoModel) error {
	return fmt.Errorf("database error")
}

func TestSavePhoto_Happy(t *testing.T) {
	routes.SetPhotoService(mockAddPhotoSuccess)

	r := gin.Default()
	r.POST("/telemetry/photo", routes.SavePhoto)

	photo := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAABnSURBVDhP7cyhDcAgDERRZ5fMkhmYgQ1S0NLR0dHQUdHxJQcJpEBCKv7kSJZ1W3d5wQvGmBzH4TiO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4/wBZ8kQCPuQ6VQAAAAASUVORK5CYII=",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	jsonPhoto, _ := json.Marshal(photo)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/photo", bytes.NewBuffer(jsonPhoto))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, but got %v", http.StatusCreated, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["status"] != "Photo Saved Successfully" {
		t.Errorf("Expected success message 'Photo Saved Successfully', but got %v", response["status"])
	}
}

func TestSavePhoto_BadRequest(t *testing.T) {
	routes.SetPhotoService(mockAddPhotoSuccess)

	r := gin.Default()
	r.POST("/telemetry/photo", routes.SavePhoto)

	photo := map[string]interface{}{"bad": "bad"}

	jsonPhoto, _ := json.Marshal(photo)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/photo", bytes.NewBuffer(jsonPhoto))
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

func TestSavePhoto_InternalServerError(t *testing.T) {
	routes.SetPhotoService(mockAddPhotoError)

	r := gin.Default()
	r.POST("/telemetry/photo", routes.SavePhoto)

	photo := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAABnSURBVDhP7cyhDcAgDERRZ5fMkhmYgQ1S0NLR0dHQUdHxJQcJpEBCKv7kSJZ1W3d5wQvGmBzH4TiO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4/wBZ8kQCPuQ6VQAAAAASUVORK5CYII=",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	jsonPhoto, _ := json.Marshal(photo)

	req, _ := http.NewRequest(http.MethodPost, "/telemetry/photo", bytes.NewBuffer(jsonPhoto))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %v, but got %v", http.StatusInternalServerError, recorder.Code)
	}

	var response map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["error"] != "Error saving photo to database: database error" {
		t.Errorf("Expected error message 'Error saving photo to database: database error', but got %v", response["error"])
	}
}
