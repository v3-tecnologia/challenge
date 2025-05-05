package entity_test

import (
	"testing"
	"time"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewGyroscope_ValidData(t *testing.T) {
	gyro, err := entity.NewGyroscope("00:11:22:33:44:55", 1.2, 3.4, 5.6, time.Now())
	assert.NotNil(t, gyro)
	assert.NoError(t, err)
}

func TestNewGyroscope_MissingDeviceID(t *testing.T) {
	gyro, err := entity.NewGyroscope("", 1.0, 2.0, 3.0, time.Now())
	assert.Nil(t, gyro)
	assert.EqualError(t, err, "deviceID cannot be empty")
}

func TestNewGyroscope_ZeroTimestamp(t *testing.T) {
	gyro, err := entity.NewGyroscope("device-123", 1.0, 2.0, 3.0, time.Time{})
	assert.Nil(t, gyro)
	assert.EqualError(t, err, "timestamp is required")
}
