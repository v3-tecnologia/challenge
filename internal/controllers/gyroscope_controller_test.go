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

type MockGyroscopeService struct {
	mock.Mock
}

func (m *MockGyroscopeService) Save(data *models.Gyroscope) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockGyroscopeService) GetAll(page, size int) ([]models.Gyroscope, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.Gyroscope), args.Error(1)
}

func TestCreateGyroscope_Success(t *testing.T) {
	mockService := new(MockGyroscopeService)
	controller := controller.NewGyroscopeController(mockService)

	payload := models.Gyroscope{
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
		MAC:       "AA:BB:CC:DD:EE:FF",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mockService.On("Save", mock.AnythingOfType("*models.Gyroscope")).Return(nil)

	controller.CreateGyroscope(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestCreateGyroscope_BadRequest(t *testing.T) {
	mockService := new(MockGyroscopeService)
	controller := controller.NewGyroscopeController(mockService)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBufferString(`{invalid json}`))
	rec := httptest.NewRecorder()

	controller.CreateGyroscope(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateGyroscope_BadRequest_JSON_Response(t *testing.T) {
	mockService := new(MockGyroscopeService)
	controller := controller.NewGyroscopeController(mockService)

	tests := []struct {
		name        string
		payload     models.Gyroscope
		expectedMsg map[string]interface{}
	}{
		{
			name:    "Missing X",
			payload: models.Gyroscope{Y: 1.1, Z: 2.2, MAC: "AA:BB:CC:DD:EE:FF"},
			expectedMsg: map[string]interface{}{
				"X": "Field axis X is require",
			},
		},
		{
			name:    "Missing Y",
			payload: models.Gyroscope{X: 1.1, Z: 2.2, MAC: "AA:BB:CC:DD:EE:FF"},

			expectedMsg: map[string]interface{}{
				"Y": "Field axis Y is require",
			},
		},
		{
			name:    "Missing Z",
			payload: models.Gyroscope{X: 1.1, Y: 2.2, MAC: "AA:BB:CC:DD:EE:FF"},

			expectedMsg: map[string]interface{}{
				"Z": "Field axis Z is require",
			},
		},
		{
			name:    "Missing MAC",
			payload: models.Gyroscope{X: 1.1, Y: 2.2, Z: 3.3},
			expectedMsg: map[string]interface{}{
				"MAC": "Field MAC is require",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			controller.CreateGyroscope(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			var resp map[string]interface{}
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMsg, resp["message"])
		})
	}
}
