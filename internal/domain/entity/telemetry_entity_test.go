package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGyroscopeValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   GyroscopeTelemetry
		isValid bool
	}{
		{
			name: "Valid Gyroscope",
			input: GyroscopeTelemetry{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "00:11:22:33:44:55",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
			},
			isValid: true,
		},
		{
			name: "Invalid Gyroscope - Empty DeviceID",
			input: GyroscopeTelemetry{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				X: 1.0,
				Y: 2.0,
				Z: 3.0,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGPSValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   GPSTelemetry
		isValid bool
	}{
		{
			name: "Valid GPS",
			input: GPSTelemetry{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "00:11:22:33:44:55",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				Latitude:  -23.550520,
				Longitude: -46.633308,
			},
			isValid: true,
		},
		{
			name: "Invalid GPS - Invalid Latitude",
			input: GPSTelemetry{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "00:11:22:33:44:55",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				Latitude:  91.0,
				Longitude: -46.633308,
			},
			isValid: false,
		},
		{
			name: "Invalid GPS - Empty DeviceID",
			input: GPSTelemetry{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				Latitude:  -23.550520,
				Longitude: -46.633308,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestPhotoValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   Picture
		isValid bool
	}{
		{
			name: "Valid Photo",
			input: Picture{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "00:11:22:33:44:55",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				PictureURL: "https://example.com/image.jpg",
			},
			isValid: true,
		},
		{
			name: "Invalid Photo - Empty ImageURL",
			input: Picture{
				BaseEntity: BaseEntity{
					ID:         uuid.New(),
					DeviceID:   "00:11:22:33:44:55",
					CreatedAt:  time.Now(),
					ReceivedAt: time.Now(),
				},
				PictureURL: "",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
