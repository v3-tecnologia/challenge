package entity_test

import (
	"testing"
	"time"

	"github.com/mkafonso/go-cloud-challenge/entity"
	"github.com/mkafonso/go-cloud-challenge/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewGPS_ValidData(t *testing.T) {
	gps, err := entity.NewGPS("00:11:22:33:44:55", utils.Ptr(10.0), utils.Ptr(10.0), time.Now())
	assert.NotNil(t, gps)
	assert.NoError(t, err)
}

func TestNewGPS_EmptyDeviceID(t *testing.T) {
	gps, err := entity.NewGPS("", utils.Ptr(10.0), utils.Ptr(10.0), time.Now())
	assert.Nil(t, gps)
	assert.EqualError(t, err, "device_id is required")
}

func TestNewGPS_EmptyTimestamp(t *testing.T) {
	gps, err := entity.NewGPS("00:11:22:33:44:55", utils.Ptr(10.0), utils.Ptr(10.0), time.Time{})
	assert.Nil(t, gps)
	assert.EqualError(t, err, "timestamp is required")
}
