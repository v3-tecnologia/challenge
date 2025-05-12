package domain

import (
	"errors"
	"math"
	"testing"
	"time"
	"v3/internal/domain"
)

func TestNewGPSData(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.GPSDto
		wantErr error
	}{
		{
			name: "Valid GPS data",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			wantErr: nil,
		},
		{
			name: "Invalid Device ID",
			input: &domain.GPSDto{
				DeviceID:  "invalid-mac",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			wantErr: domain.ErrDeviceIDGPS,
		},
		{
			name: "Zero timestamp",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			wantErr: domain.ErrTimestampGPS,
		},
		{
			name: "Invalid Latitude (NaN)",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  math.NaN(),
				Longitude: -74.0060,
			},
			wantErr: domain.ErrInvalidGPSValues,
		},
		{
			name: "Invalid Longitude (Inf)",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: math.Inf(1),
			},
			wantErr: domain.ErrInvalidGPSValues,
		},
		{
			name: "Invalid Latitude out of range",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  100.0, // Latitude inválida, maior que 90
				Longitude: -74.0060,
			},
			wantErr: domain.ErrInvalidGPSValues,
		},
		{
			name: "Invalid Longitude out of range",
			input: &domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				Latitude:  40.7128,
				Longitude: 200.0, // Longitude inválida, maior que 180
			},
			wantErr: domain.ErrInvalidGPSValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewGPSData(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewGPSData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got == nil {
				t.Error("Expected GPS data, got nil")
			}
		})
	}
}
