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

type MockPhotoService struct {
	mock.Mock
}

func (m *MockPhotoService) Save(data *models.Photo) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockPhotoService) GetAll(page, size int) ([]models.Photo, error) {
	args := m.Called(page, size)
	return args.Get(0).([]models.Photo), args.Error(1)
}

func TestCreatePhoto_Success(t *testing.T) {
	mockService := new(MockPhotoService)
	controller := controller.NewPhotoController(mockService)

	payload := models.Photo{
		ImageURL:   "",
		Recognized: false,
		MAC:        "AA:BB:CC:DD:EE:FF",
		Timestamp:  time.Now(),
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mockService.On("Save", mock.AnythingOfType("*models.Photo")).Return(nil)

	controller.CreatePhoto(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestCreatePhoto_BadRequest(t *testing.T) {
	mockService := new(MockPhotoService)
	controller := controller.NewPhotoController(mockService)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBufferString(`{invalid json}`))
	rec := httptest.NewRecorder()

	controller.CreatePhoto(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreatePhoto_BadRequest_JSON_Response(t *testing.T) {
	mockService := new(MockPhotoService)
	controller := controller.NewPhotoController(mockService)

	tests := []struct {
		name          string
		payload       models.Photo
		expectedError map[string]interface{}
		expectedMsg   map[string]interface{}
	}{

		{
			name:    "Missing MAC",
			payload: models.Photo{ImageURL: ""},
			expectedError: map[string]interface{}{
				"MAC": "Key: 'Photo.MAC' Error:Field validation for 'MAC' failed on the 'required' tag",
			},
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

			controller.CreatePhoto(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)

			var resp map[string]interface{}
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedError, resp["error"])
			assert.Equal(t, tt.expectedMsg, resp["message"])
		})
	}
}
