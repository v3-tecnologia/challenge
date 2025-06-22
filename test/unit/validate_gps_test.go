package unit

import (
	"testing"

	"github.com/yanvic/challenge/core/entity"
	"github.com/yanvic/challenge/core/usecase"
)

func TestValidateGPS(t *testing.T) {
	lat := -23.55
	long := -46.63

	tests := []struct {
		name    string
		input   entity.GPS
		wantErr bool
	}{
		{
			name: "valid data",
			input: entity.GPS{
				Latitude:  &lat,
				Longitude: &long,
				Timestamp: "2025-06-21T15:00:00Z",
				DeviceID:  "device-123",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := usecase.ValidateGPS(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGPS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
