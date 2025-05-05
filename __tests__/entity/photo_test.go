package entity_test

import (
	"testing"
	"time"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewPhoto_ValidData(t *testing.T) {
	photo, err := entity.NewPhoto("00:11:22:33:44:55", "/photos/img.jpg", time.Now(), true)
	assert.NotNil(t, photo)
	assert.NoError(t, err)
}

func TestNewPhoto_EmptyDeviceID(t *testing.T) {
	photo, err := entity.NewPhoto("", "/photos/img.jpg", time.Now(), false)
	assert.Nil(t, photo)
	assert.EqualError(t, err, "device_id is required")
}

func TestNewPhoto_EmptyFilePath(t *testing.T) {
	photo, err := entity.NewPhoto("device-123", "", time.Now(), false)
	assert.Nil(t, photo)
	assert.EqualError(t, err, "file_path is required")
}

func TestNewPhoto_ZeroTimestamp(t *testing.T) {
	photo, err := entity.NewPhoto("device-123", "/photos/img.jpg", time.Time{}, false)
	assert.Nil(t, photo)
	assert.EqualError(t, err, "timestamp is required")
}
