package usecase_test

import (
	"image"
	"testing"

	"bytes"
	"image/jpeg"

	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/core/usecase"
)

func TestValidateGPS(t *testing.T) {
	lat := 1.0
	lon := 2.0

	tests := []struct {
		name    string
		input   entity.GPS
		wantErr string
	}{
		{
			name: "valid input",
			input: entity.GPS{
				Latitude:  &lat,
				Longitude: &lon,
				DeviceID:  "dev123",
				Timestamp: "2025-06-21T15:00:00Z",
			},
			wantErr: "",
		},
		{
			name:    "missing device_id",
			input:   entity.GPS{Latitude: &lat, Longitude: &lon, Timestamp: "2025-06-21"},
			wantErr: "device_id is required",
		},
		{
			name:    "missing timestamp",
			input:   entity.GPS{Latitude: &lat, Longitude: &lon, DeviceID: "dev123"},
			wantErr: "timestamp is required",
		},
		{
			name:    "nil latitude",
			input:   entity.GPS{Longitude: &lon, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "latitude is required",
		},
		{
			name:    "nil longitude",
			input:   entity.GPS{Latitude: &lat, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "longitude is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.ValidateGPS(tt.input)
			if tt.wantErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tt.wantErr != "" && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("expected error: %s, got: %v", tt.wantErr, err)
			}
		})
	}
}

func TestValidateGyroscope(t *testing.T) {
	x, y, z := 1.0, 2.0, 3.0

	tests := []struct {
		name    string
		input   entity.Gyroscope
		wantErr string
	}{
		{
			name: "valid input",
			input: entity.Gyroscope{
				X:         &x,
				Y:         &y,
				Z:         &z,
				DeviceID:  "dev123",
				Timestamp: "2025-06-21",
			},
			wantErr: "",
		},
		{
			name:    "missing device_id",
			input:   entity.Gyroscope{X: &x, Y: &y, Z: &z, Timestamp: "2025-06-21"},
			wantErr: "device_id is required",
		},
		{
			name:    "missing timestamp",
			input:   entity.Gyroscope{X: &x, Y: &y, Z: &z, DeviceID: "dev123"},
			wantErr: "timestamp is required",
		},
		{
			name:    "nil x",
			input:   entity.Gyroscope{Y: &y, Z: &z, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "x is required",
		},
		{
			name:    "nil y",
			input:   entity.Gyroscope{X: &x, Z: &z, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "y is required",
		},
		{
			name:    "nil z",
			input:   entity.Gyroscope{X: &x, Y: &y, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "z is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.ValidateGyroscope(tt.input)
			if tt.wantErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tt.wantErr != "" && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("expected error: %s, got: %v", tt.wantErr, err)
			}
		})
	}
}

func TestValidatePhoto(t *testing.T) {
	var buf bytes.Buffer
	_ = jpeg.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 1, 1)), nil)
	validImage := buf.Bytes()

	tests := []struct {
		name    string
		input   entity.Photo
		wantErr string
	}{
		{
			name: "valid input",
			input: entity.Photo{
				Image:     validImage,
				DeviceID:  "dev123",
				Timestamp: "2025-06-21",
			},
			wantErr: "",
		},
		{
			name:    "missing device_id",
			input:   entity.Photo{Image: validImage, Timestamp: "2025-06-21"},
			wantErr: "device_id is required",
		},
		{
			name:    "missing timestamp",
			input:   entity.Photo{Image: validImage, DeviceID: "dev123"},
			wantErr: "timestamp is required",
		},
		{
			name:    "empty image",
			input:   entity.Photo{Image: []byte{}, DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "image data is required",
		},
		{
			name:    "invalid image format",
			input:   entity.Photo{Image: []byte("not-an-image"), DeviceID: "dev123", Timestamp: "2025-06-21"},
			wantErr: "image data is not a valid image format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.ValidatePhoto(tt.input)
			if tt.wantErr == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if tt.wantErr != "" && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("expected error: %s, got: %v", tt.wantErr, err)
			}
		})
	}
}
