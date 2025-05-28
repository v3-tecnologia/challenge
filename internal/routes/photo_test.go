package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockPhotoService struct {
	ErrToReturn error
	Recognized  bool
}

func (m *MockPhotoService) AddPhoto(photo models.PhotoModel) (bool, error) {
	return m.Recognized, m.ErrToReturn
}

func TestSavePhoto_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPhotoService{Recognized: true}
	handler := routes.NewPhotoHandler(mockService)

	router := gin.Default()
	router.POST("/photo", handler.SavePhoto)

	payload := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg==",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Photo Saved Successfully")
	assert.Contains(t, w.Body.String(), "recognized")
}

func TestSavePhoto_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPhotoService{}
	handler := routes.NewPhotoHandler(mockService)

	router := gin.Default()
	router.POST("/photo", handler.SavePhoto)

	req := httptest.NewRequest(http.MethodPost, "/photo", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSavePhoto_ValidationErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPhotoService{}
	handler := routes.NewPhotoHandler(mockService)

	router := gin.Default()
	router.POST("/photo", handler.SavePhoto)

	req := httptest.NewRequest(http.MethodPost, "/photo", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "errors")
}

func TestSavePhoto_InternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPhotoService{
		ErrToReturn: errors.New("unexpected error"),
	}
	handler := routes.NewPhotoHandler(mockService)

	router := gin.Default()
	router.POST("/photo", handler.SavePhoto)

	payload := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg==",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error processing photo")
}
