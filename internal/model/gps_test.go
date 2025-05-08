package model

import (
	"testing"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGPSBuilder_WithMacAddress(t *testing.T) {
	tests := []struct {
		macAddress string
		errCode    string
		errMessage string
	}{
		{"00:14:22:01:23:45", "", ""},
		{"", "missing_device_id", "device identifier (MAC address) is required"},
	}

	for _, tt := range tests {
		t.Run(tt.macAddress, func(t *testing.T) {
			builder := NewGPSBuilder().WithMacAddress(tt.macAddress)
			if tt.errCode != "" {
				assert.NotNil(t, builder.err)
				assert.Equal(t, tt.errCode, builder.err.(*errors.AppError).Code)
				assert.Equal(t, tt.errMessage, builder.err.(*errors.AppError).Message)
			} else {
				assert.Nil(t, builder.err)
			}
		})
	}
}

func TestGPSBuilder_WithCoordinates(t *testing.T) {
	tests := []struct {
		latitude   *float64
		longitude  *float64
		errCode    string
		errMessage string
	}{
		{floatPointer(-46.633308), floatPointer(-23.55052), "", ""},
		{nil, floatPointer(-23.55052), "missing_latitude", "latitude coordinate is required"},
		{floatPointer(-46.633308), nil, "missing_longitude", "longitude coordinate is required"},
	}

	for _, tt := range tests {
		t.Run("Coordinates", func(t *testing.T) {
			builder := NewGPSBuilder().WithCoordinates(tt.latitude, tt.longitude)
			if tt.errCode != "" {
				assert.NotNil(t, builder.err)
				assert.Equal(t, tt.errCode, builder.err.(*errors.AppError).Code)
				assert.Equal(t, tt.errMessage, builder.err.(*errors.AppError).Message)
			} else {
				assert.Nil(t, builder.err)
			}
		})
	}
}

func TestGPSBuilder_WithTimestamp(t *testing.T) {
	validTimestamp := time.Now()
	invalidTimestamp := time.Time{}

	tests := []struct {
		timestamp  time.Time
		errCode    string
		errMessage string
	}{
		{validTimestamp, "", ""},
		{invalidTimestamp, "invalid_timestamp", "timestamp is required and must be valid"},
	}

	for _, tt := range tests {
		t.Run("Timestamp", func(t *testing.T) {
			builder := NewGPSBuilder().WithTimestamp(tt.timestamp)
			if tt.errCode != "" {
				assert.NotNil(t, builder.err)
				assert.Equal(t, tt.errCode, builder.err.(*errors.AppError).Code)
				assert.Equal(t, tt.errMessage, builder.err.(*errors.AppError).Message)
			} else {
				assert.Nil(t, builder.err)
			}
		})
	}
}

func TestGPSBuilder_Build(t *testing.T) {
	validLatitude := float64(-46.633308)
	validLongitude := float64(-23.55052)
	validTimestamp := time.Now()

	tests := []struct {
		macAddress string
		latitude   *float64
		longitude  *float64
		timestamp  time.Time
		expectErr  bool
	}{
		{"00:14:22:01:23:45", &validLatitude, &validLongitude, validTimestamp, false}, // Caso válido
		{"", &validLatitude, &validLongitude, validTimestamp, true},                   // MacAddress vazio
		{"00:14:22:01:23:45", nil, &validLongitude, validTimestamp, true},             // Latitude nula
		{"00:14:22:01:23:45", &validLatitude, nil, validTimestamp, true},              // Longitude nula
		{"00:14:22:01:23:45", &validLatitude, &validLongitude, time.Time{}, true},     // Timestamp inválido
	}

	for _, tt := range tests {
		t.Run(tt.macAddress, func(t *testing.T) {
			builder := NewGPSBuilder().
				WithMacAddress(tt.macAddress).
				WithCoordinates(tt.latitude, tt.longitude).
				WithTimestamp(tt.timestamp)

			gps, err := builder.Build()

			if tt.expectErr {
				assert.NotNil(t, err)
				assert.Nil(t, gps)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, gps)
				assert.Equal(t, tt.macAddress, gps.MacAddress)
				assert.Equal(t, *tt.latitude, *gps.Latitude)
				assert.Equal(t, *tt.longitude, *gps.Longitude)
				assert.Equal(t, tt.timestamp, gps.Timestamp)
				assert.NotEqual(t, uuid.Nil, gps.ID)
			}
		})
	}
}

func floatPointer(f float64) *float64 {
	return &f
}
