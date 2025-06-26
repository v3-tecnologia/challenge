package telemetriesControllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"v3-test/internal/dtos/telemetriesDtos"
	"v3-test/internal/models/telemetriesModels"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockGpsUsecase struct {
	mockCreateFunc func(dto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error)
}

func (m *mockGpsUsecase) CreateGps(gpsDto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error) {
	return m.mockCreateFunc(gpsDto)
}

func setupGpsRouters(controller GpsController) *gin.Engine {
	router := gin.Default()
	router.POST("/telemetry/gps", controller.CreateGps)
	return router
}

func TestCreateGps_Success(t *testing.T) {
	mockUsecase := &mockGpsUsecase{
		mockCreateFunc: func(dto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error) {
			return telemetriesModels.GpsModel{
				Latitude:  dto.Latitude,
				Longitude: dto.Longitude,
			}, nil
		},
	}
	controller := NewGpsController(mockUsecase)
	router := setupGpsRouters(controller)

	payload := telemetriesDtos.CreateGpsDto{
		Latitude:  F64(1.23),
		Longitude: F64(4.56),
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"data"`)
}

func TestCreateGps_InvalidBody(t *testing.T) {
	mockUsecase := &mockGpsUsecase{}
	controller := NewGpsController(mockUsecase)
	router := setupGpsRouters(controller)

	invalidJSON := `{ "Latitude": "invalid" }`

	req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}

func TestErrorOnGpsCreateUsecaseCreate(t *testing.T) {
	mockUsecase := &mockGpsUsecase{
		mockCreateFunc: func(dto telemetriesDtos.CreateGpsDto) (telemetriesModels.GpsModel, error) {
			return telemetriesModels.GpsModel{}, errors.New("failed to create gps")
		},
	}

	controller := NewGpsController(mockUsecase)
	router := setupGpsRouters(controller)

	payload := telemetriesDtos.CreateGpsDto{
		Latitude:  F64(1.23),
		Longitude: F64(4.56),
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	assert.Contains(t, w.Body.String(), `"error"`)
}
