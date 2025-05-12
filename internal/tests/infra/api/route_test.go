package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"v3/internal/domain"
	"v3/internal/infra/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockCreateGyroscopeUseCase is a mock implementation of the use case for testing
type MockCreateGyroscopeUseCase struct{}

func (m *MockCreateGyroscopeUseCase) Execute(input domain.GyroscopeDto) (*domain.Gyroscope, error) {
	// Simulate successful creation of a gyroscope object
	return &domain.Gyroscope{
		ID:        "1",
		DeviceID:  input.DeviceID,
		Timestamp: input.Timestamp,
	}, nil
}

// MockGPSHandlers struct
type MockGPSHandlers struct{}

func (m *MockGPSHandlers) SetupRoutes(router *gin.Engine) {
	router.POST("/api/telemetry/gps", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GPS route"})
	})
}

// MockPhotoHandlers struct
type MockPhotoHandlers struct{}

func (m *MockPhotoHandlers) SetupRoutes(router *gin.Engine) {
	router.POST("/api/telemetry/photo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Photo route"})
	})
}

// Test function for SetupRouter
func TestSetupRouter(t *testing.T) {
	// Create the mocks for the handlers
	gyroscopeHandlers := &api.GyroscopeHandlers{
		CreateGyroscopeUseCase: &MockCreateGyroscopeUseCase{},
	}
	gpsHandlers := &MockGPSHandlers{}
	photoHandlers := &MockPhotoHandlers{}

	// Call SetupRouter
	router := api.SetupRouter(gyroscopeHandlers, gpsHandlers, photoHandlers)

	// Test the /api/telemetry/gyroscope route
	w := performRequest(router, http.MethodPost, "/api/telemetry/gyroscope", `{"device_id": "123", "timestamp": "2025-01-01T00:00:00Z"}`)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"ID": "1", "DeviceID": "123", "Timestamp": "2025-01-01T00:00:00Z"}`, w.Body.String())

	// Test the /api/telemetry/gps route
	w = performRequest(router, http.MethodPost, "/api/telemetry/gps", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "GPS route"}`, w.Body.String())

	// Test the /api/telemetry/photo route
	w = performRequest(router, http.MethodPost, "/api/telemetry/photo", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Photo route"}`, w.Body.String())
}

// Helper function to perform HTTP requests
func performRequest(router *gin.Engine, method, url string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, bytes.NewBufferString(body)) // use body here if needed
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
