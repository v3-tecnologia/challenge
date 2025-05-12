package device_test

import (
	"errors"
	"testing"
	"v3/internal/domain"
)

func TestNewDevice(t *testing.T) {
	tests := []struct {
		name    string
		mac     string
		wantID  string
		wantErr error
	}{
		{
			name:    "Valid MAC address",
			mac:     "00:0a:95:9d:68:16",
			wantID:  "00:0a:95:9d:68:16",
			wantErr: nil,
		},
		{
			name:    "Valid MAC with uppercase",
			mac:     "FF:FF:FF:FF:FF:FF",
			wantID:  "FF:FF:FF:FF:FF:FF",
			wantErr: nil,
		},
		{
			name:    "Invalid MAC - too short",
			mac:     "00:0a:95:9d:68",
			wantID:  "",
			wantErr: domain.ErrInValidMACGyroscope,
		},
		{
			name:    "Invalid MAC - wrong format",
			mac:     "00-0a-95-9d-68-16",
			wantID:  "",
			wantErr: domain.ErrInValidMACGyroscope,
		},
		{
			name:    "Invalid MAC - non-hex characters",
			mac:     "00:0a:95:9d:68:GG",
			wantID:  "",
			wantErr: domain.ErrInValidMACGyroscope,
		},
		{
			name:    "Invalid MAC - empty string",
			mac:     "",
			wantID:  "",
			wantErr: domain.ErrInValidMACGyroscope,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewDevice(tt.mac)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewDevice(%q) error = %v, wantErr %v", tt.mac, err, tt.wantErr)
				return
			}
			if err == nil {
				if got == nil || got.ID != tt.wantID {
					t.Errorf("NewDevice(%q) = %v, want ID %q", tt.mac, got, tt.wantID)
				}
			}
		})
	}
}
