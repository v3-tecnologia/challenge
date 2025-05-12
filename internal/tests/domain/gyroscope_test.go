package domain

import (
	"errors"
	"math"
	"testing"
	"time"
	"v3/internal/domain"
)

func TestNewGyroscopeData(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.GyroscopeDto
		wantErr error
	}{
		{
			name: "Valid gyroscope data",
			input: &domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.0,
				Y:         2.0,
				Z:         3.0,
			},
			wantErr: nil,
		},
		{
			name: "Invalid Device ID",
			input: &domain.GyroscopeDto{
				DeviceID:  "invalid-mac",
				Timestamp: time.Now().Unix(),
				X:         1.0,
				Y:         2.0,
				Z:         3.0,
			},
			wantErr: domain.ErrDeviceIDGyroscope,
		},
		{
			name: "Zero timestamp",
			input: &domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
				X:         1.0,
				Y:         2.0,
				Z:         3.0,
			},
			wantErr: nil,
		},
		{
			name: "Invalid X (NaN)",
			input: &domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         math.NaN(),
				Y:         2.0,
				Z:         3.0,
			},
			wantErr: domain.ErrInvalidGyroscopeValues,
		},
		{
			name: "Invalid Y (Inf)",
			input: &domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.0,
				Y:         math.Inf(1),
				Z:         3.0,
			},
			wantErr: domain.ErrInvalidGyroscopeValues,
		},
		{
			name: "Invalid Z (NaN)",
			input: &domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
				X:         1.0,
				Y:         2.0,
				Z:         math.NaN(),
			},
			wantErr: domain.ErrInvalidGyroscopeValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewGyroscopeData(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewGyroscopeData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got == nil {
				t.Error("Expected gyroscope data, got nil")
			}
		})
	}
}
