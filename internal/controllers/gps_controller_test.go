package controller_test

import (
	"bytes"
	controller "challenge-cloud/internal/controllers"
	"challenge-cloud/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGPSService struct {
	mock.Mock
}

func (m *MockGPSService) Save(data *models.GPS) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockGPSService) GetAll(page, size int) ([]models.GPS, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.GPS), args.Error(1)
}

func TestCreateGPS_Success(t *testing.T) {
	mockService := new(MockGPSService)
	controller := controller.NewGPSController(mockService)

	payload := models.GPS{
		Longitude: 1.0,
		Latitude:  2.0,
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mockService.On("Save", mock.AnythingOfType("*models.GPS")).Return(nil)

	controller.CreateGPS(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestCreateGPS_BadRequest(t *testing.T) {
	mockService := new(MockGPSService)
	controller := controller.NewGPSController(mockService)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBufferString(`{invalid json}`))
	rec := httptest.NewRecorder()

	controller.CreateGPS(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateGPS_BadRequest_JSON_Response(t *testing.T) {
	mockService := new(MockGPSService)
	controller := controller.NewGPSController(mockService)

	tests := []struct {
		name        string
		payload     models.GPS
		expectedMsg map[string]interface{}
	}{
		{
			name:    "Missing Longitude",
			payload: models.GPS{Latitude: 2.2, MAC: "AA:BB:CC:DD:EE:FF"},
			expectedMsg: map[string]interface{}{
				"Longitude": "Field Longitude is require",
			},
		},
		{
			name:    "Missing Latitude",
			payload: models.GPS{Longitude: 1.1, MAC: "AA:BB:CC:DD:EE:FF"},
			expectedMsg: map[string]interface{}{
				"Latitude": "Field Latitude is require",
			},
		},
		{
			name:    "Missing MAC",
			payload: models.GPS{Latitude: 1.1, Longitude: 2.2},
			expectedMsg: map[string]interface{}{
				"MAC": "Field MAC is require",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			controller.CreateGPS(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			var resp map[string]interface{}
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMsg, resp["message"])
		})
	}
}
