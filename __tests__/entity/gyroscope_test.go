package entity_test

import (
	"testing"
	"time"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/mkafonso/go-cloud-challenge/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewGyroscope_ValidData(t *testing.T) {
	gyro, err := entity.NewGyroscope("00:11:22:33:44:55", utils.Ptr(1.2), utils.Ptr(3.4), utils.Ptr(5.6), time.Now())
	assert.NotNil(t, gyro)
	assert.NoError(t, err)
}

func TestNewGyroscope_MissingDeviceID(t *testing.T) {
	gyro, err := entity.NewGyroscope("", utils.Ptr(1.0), utils.Ptr(2.0), utils.Ptr(3.0), time.Now())
	assert.Nil(t, gyro)
	assert.EqualError(t, err, "device_id is required")
}

func TestNewGyroscope_ZeroTimestamp(t *testing.T) {
	gyro, err := entity.NewGyroscope("device-123", utils.Ptr(1.0), utils.Ptr(2.0), utils.Ptr(3.0), time.Time{})
	assert.Nil(t, gyro)
	assert.EqualError(t, err, "timestamp is required")
}
