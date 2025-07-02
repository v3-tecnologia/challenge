package unit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of TelemetryService
type TelemetryService struct{}

func (s *TelemetryService) ProcessGyroscopeData(x, y, z float64) error {
	if x == 0 && y == 0 && z == 0 {
		return errors.New("invalid gyroscope data")
	}
	return nil
}

func (s *TelemetryService) ProcessGPSData(latitude, longitude float64) error {
	if latitude == 0 && longitude == 0 {
		return errors.New("invalid GPS data")
	}
	return nil
}

func (s *TelemetryService) ProcessPhotoData(image string) error {
	if image == "" {
		return errors.New("invalid photo data")
	}
	return nil
}

func TestProcessGyroscopeData(t *testing.T) {
	mockService := &TelemetryService{}

	// Test with valid data
	err := mockService.ProcessGyroscopeData(1.0, 2.0, 3.0)
	assert.NoError(t, err, "Expected no error for valid gyroscope data")

	// Test with invalid data (e.g., all zeros)
	err = mockService.ProcessGyroscopeData(0.0, 0.0, 0.0)
	assert.Error(t, err, "Expected error for invalid gyroscope data")
}

func TestProcessGPSData(t *testing.T) {
	mockService := &TelemetryService{}

	// Test with valid data
	err := mockService.ProcessGPSData(12.34, 56.78)
	assert.NoError(t, err, "Expected no error for valid GPS data")

	// Test with invalid data (e.g., out-of-range latitude/longitude)
	err = mockService.ProcessGPSData(0, 0) // Invalid latitude/longitude
	assert.Error(t, err, "Expected error for invalid GPS data")
}

func TestProcessPhotoData(t *testing.T) {
	mockService := &TelemetryService{}

	// Test with valid data
	err := mockService.ProcessPhotoData("base64encodedstring")
	assert.NoError(t, err, "Expected no error for valid photo data")

	// Test with invalid data (e.g., empty string)
	err = mockService.ProcessPhotoData("")
	assert.Error(t, err, "Expected error for empty photo data")
}
