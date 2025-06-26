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

type mockGyroscopeUsecase struct {
	mockCreateFunc func(dto telemetriesDtos.CreateGyroscopeDto) (telemetriesModels.GyroscopeModel, error)
}

func (m *mockGyroscopeUsecase) CreateGyroscope(dto telemetriesDtos.CreateGyroscopeDto) (telemetriesModels.GyroscopeModel, error) {
	return m.mockCreateFunc(dto)
}

func setupGyrosopeRouters(controller GyroscopeController) *gin.Engine {
	r := gin.Default()
	r.POST("/telemetry/gyroscope", controller.CreateGyroscope)
	return r
}

func TestCreateGyroscope_Success(t *testing.T) {
	mockUsecase := &mockGyroscopeUsecase{
		mockCreateFunc: func(dto telemetriesDtos.CreateGyroscopeDto) (telemetriesModels.GyroscopeModel, error) {
			return telemetriesModels.GyroscopeModel{
				X: dto.X,
				Y: dto.Y,
				Z: dto.Z,
			}, nil
		},
	}
	controller := NewGyroscopeController(mockUsecase)
	router := setupGyrosopeRouters(controller)

	payload := telemetriesDtos.CreateGyroscopeDto{
		X: F64(1.23),
		Y: F64(4.56),
		Z: F64(7.89),
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"data"`)
}

func TestCreateGyroscope_InvalidBody(t *testing.T) {
	mockUsecase := &mockGyroscopeUsecase{}
	controller := NewGyroscopeController(mockUsecase)
	router := setupGyrosopeRouters(controller)

	invalidJSON := `{ "X": "invalid" }`

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}

func TestErrorOnGyroscopeUsecaseCreate(t *testing.T) {
	mockUsecase := &mockGyroscopeUsecase{
		mockCreateFunc: func(dto telemetriesDtos.CreateGyroscopeDto) (telemetriesModels.GyroscopeModel, error) {
			return telemetriesModels.GyroscopeModel{}, errors.New("failed to create gyroscope")
		},
	}

	controller := NewGyroscopeController(mockUsecase)
	router := setupGyrosopeRouters(controller)

	payload := telemetriesDtos.CreateGyroscopeDto{
		X: F64(1.23),
		Y: F64(4.56),
		Z: F64(7.89),
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	assert.Contains(t, w.Body.String(), `"error"`)
}
