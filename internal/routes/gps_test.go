package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockGpsService struct {
	ErrToReturn error
}

func (m *MockGpsService) AddGps(gps models.GpsModel) error {
	return m.ErrToReturn
}

func TestSaveGps_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGpsService{}
	handler := routes.NewGpsHandler(mockService)

	router := gin.Default()
	router.POST("/gps", handler.SaveGps)

	payload := models.GpsModel{
		Latitude:  12.34,
		Longitude: 56.78,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "GPS Saved Successfully")
}

func TestSaveGps_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGpsService{}
	handler := routes.NewGpsHandler(mockService)

	router := gin.Default()
	router.POST("/gps", handler.SaveGps)

	req := httptest.NewRequest(http.MethodPost, "/gps", bytes.NewBuffer([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSaveGps_ValidationErrors(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGpsService{}
	handler := routes.NewGpsHandler(mockService)

	router := gin.Default()
	router.POST("/gps", handler.SaveGps)

	req := httptest.NewRequest(http.MethodPost, "/gps", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "errors")
}

func TestSaveGps_DBError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGpsService{
		ErrToReturn: custom_errors.NewDBError(errors.New("DB failure"), 500),
	}
	handler := routes.NewGpsHandler(mockService)

	router := gin.Default()
	router.POST("/gps", handler.SaveGps)

	payload := models.GpsModel{
		Latitude:  12.34,
		Longitude: 56.78,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Contains(t, w.Body.String(), "A database error occurred")
}

func TestSaveGps_UnexpectedError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockGpsService{
		ErrToReturn: errors.New("something went wrong"),
	}
	handler := routes.NewGpsHandler(mockService)

	router := gin.Default()
	router.POST("/gps", handler.SaveGps)

	payload := models.GpsModel{
		Latitude:  12.34,
		Longitude: 56.78,
		MAC:       "8C:16:45:8D:F3:7B",
		Timestamp: 1746110367207,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "An unexpected server error occurred")
}
