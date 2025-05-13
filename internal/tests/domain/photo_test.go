package domain

import (
	"errors"
	"testing"
	"time"
	"v3/internal/domain"
)

func TestNewPhotoData(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.PhotoDto
		wantErr error
	}{
		{
			name: "Valid Photo data",
			input: &domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: time.Now().Unix(),
			},
			wantErr: nil,
		},
		{
			name: "Invalid Device ID",
			input: &domain.PhotoDto{
				DeviceID:  "",
				Timestamp: time.Now().Unix(),
			},
			wantErr: domain.ErrDeviceIDPhoto,
		},
		{
			name: "Zero timestamp",
			input: &domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
			},
			wantErr: domain.ErrTimestampPhoto,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewPhotoData(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewPhotoData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got == nil {
				t.Error("Expected Photo data, got nil")
			}
		})
	}
}
