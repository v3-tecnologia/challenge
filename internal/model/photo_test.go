package model

import (
	"testing"
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPhotoBuilder_WithMacAddress(t *testing.T) {
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
			builder := NewPhotoBuilder().WithMacAddress(tt.macAddress)
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

func TestPhotoBuilder_WithFileUrl(t *testing.T) {
	tests := []struct {
		fileUrl    string
		errCode    string
		errMessage string
	}{
		{"https://example.com/photo.jpg", "", ""},
		{"", "missing_file_path", "file path for the photo is required"},
	}

	for _, tt := range tests {
		t.Run(tt.fileUrl, func(t *testing.T) {
			builder := NewPhotoBuilder().WithFileUrl(tt.fileUrl)
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

func TestPhotoBuilder_WithRecognitionStatus(t *testing.T) {
	tests := []struct {
		isMatch bool
	}{
		{true},
		{false},
	}

	for _, tt := range tests {
		t.Run("WithRecognitionStatus", func(t *testing.T) {
			builder := NewPhotoBuilder().WithRecognitionStatus(tt.isMatch)
			assert.Nil(t, builder.err)
			photo, err := builder.Build()
			assert.Nil(t, err)
			assert.Equal(t, tt.isMatch, photo.IsMatch)
		})
	}
}

func TestPhotoBuilder_WithTimestamp(t *testing.T) {
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
			builder := NewPhotoBuilder().WithTimestamp(tt.timestamp)
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

func TestPhotoBuilder_Build(t *testing.T) {
	validTimestamp := time.Now()

	tests := []struct {
		macAddress string
		fileUrl    string
		isMatch    bool
		timestamp  time.Time
		expectErr  bool
	}{
		{"00:14:22:01:23:45", "https://example.com/photo.jpg", true, validTimestamp, false}, // Caso válido
		{"", "https://example.com/photo.jpg", true, validTimestamp, true},                   // MacAddress vazio
		{"00:14:22:01:23:45", "", true, validTimestamp, true},                               // FileUrl vazio
		{"00:14:22:01:23:45", "https://example.com/photo.jpg", true, time.Time{}, true},     // Timestamp inválido
	}

	for _, tt := range tests {
		t.Run(tt.macAddress, func(t *testing.T) {
			builder := NewPhotoBuilder().
				WithMacAddress(tt.macAddress).
				WithFileUrl(tt.fileUrl).
				WithRecognitionStatus(tt.isMatch).
				WithTimestamp(tt.timestamp)

			photo, err := builder.Build()

			if tt.expectErr {
				assert.NotNil(t, err)
				assert.Nil(t, photo)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, photo)
				assert.Equal(t, tt.macAddress, photo.MacAddress)
				assert.Equal(t, tt.fileUrl, photo.FileURL)
				assert.Equal(t, tt.isMatch, photo.IsMatch)
				assert.Equal(t, tt.timestamp, photo.Timestamp)
				assert.NotEqual(t, uuid.Nil, photo.ID)
			}
		})
	}
}
