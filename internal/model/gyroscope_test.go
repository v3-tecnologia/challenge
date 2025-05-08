package model

import (
	"testing"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGyroscopeBuilder_WithMacAddress(t *testing.T) {
	tests := []struct {
		macAddress string
		errCode    string
		errMessage string
	}{
		{"00:14:22:01:23:45", "", ""},
		{"", "missing_device_id", "device identifier is required"},
	}

	for _, tt := range tests {
		t.Run(tt.macAddress, func(t *testing.T) {
			builder := NewGyroscopeBuilder().WithMacAddress(tt.macAddress)
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

func TestGyroscopeBuilder_WithAxisValues(t *testing.T) {
	tests := []struct {
		x, y, z    float64
		errCode    string
		errMessage string
	}{
		{1.0, 0.0, 0.0, "", ""}, // Válido
		{0.0, 0.0, 0.0, "invalid_axis_values", "at least one axis must have a non-zero value"}, // Inválido
		{0.0, 1.0, 0.0, "", ""}, // Válido
		{0.0, 0.0, 1.0, "", ""}, // Válido
	}

	for _, tt := range tests {
		t.Run("AxisValues", func(t *testing.T) {
			builder := NewGyroscopeBuilder().WithAxisValues(tt.x, tt.y, tt.z)
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

func TestGyroscopeBuilder_WithTimestamp(t *testing.T) {
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
			builder := NewGyroscopeBuilder().WithTimestamp(tt.timestamp)
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

func TestGyroscopeBuilder_Build(t *testing.T) {
	validTimestamp := time.Now()

	tests := []struct {
		macAddress string
		x, y, z    float64
		timestamp  time.Time
		expectErr  bool
	}{
		{"00:14:22:01:23:45", 1.0, 0.0, 0.0, validTimestamp, false}, // Caso válido
		{"", 1.0, 0.0, 0.0, validTimestamp, true},                   // MacAddress vazio
		{"00:14:22:01:23:45", 0.0, 0.0, 0.0, validTimestamp, true},  // Axis vazio
		{"00:14:22:01:23:45", 1.0, 1.0, 1.0, time.Time{}, true},     // Timestamp inválido
	}

	for _, tt := range tests {
		t.Run(tt.macAddress, func(t *testing.T) {
			builder := NewGyroscopeBuilder().
				WithMacAddress(tt.macAddress).
				WithAxisValues(tt.x, tt.y, tt.z).
				WithTimestamp(tt.timestamp)

			gyroscope, err := builder.Build()

			if tt.expectErr {
				assert.NotNil(t, err)
				assert.Nil(t, gyroscope)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, gyroscope)
				assert.Equal(t, tt.macAddress, gyroscope.MacAddress)
				assert.Equal(t, tt.x, gyroscope.AxisX)
				assert.Equal(t, tt.y, gyroscope.AxisY)
				assert.Equal(t, tt.z, gyroscope.AxisZ)
				assert.Equal(t, tt.timestamp, gyroscope.Timestamp)
				assert.NotEqual(t, uuid.Nil, gyroscope.ID)
			}
		})
	}
}
