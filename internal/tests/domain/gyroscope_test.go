package gyroscope

import (
	"testing"
	"time"

	"github.com/iamrosada0/v3/internal/domain/gyroscope"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestNewGyroscopeData_Success(t *testing.T) {
	deviceID := "00:1A:2B:3C:4D:5E"
	x := gofakeit.Float64()
	y := gofakeit.Float64()
	z := gofakeit.Float64()
	timestamp := time.Now()

	data, err := gyroscope.NewGyroscopeData(deviceID, x, y, z, timestamp)
	assert.NoError(t, err)
	assert.Equal(t, deviceID, data.Device.ID)
	assert.Equal(t, x, data.X)
	assert.Equal(t, y, data.Y)
	assert.Equal(t, z, data.Z)
	assert.Equal(t, timestamp, data.Timestamp)
}
func TestNewGyroscopeData_InvalidDeviceID(t *testing.T) {
	data, err := gyroscope.NewGyroscopeData("invalid-mac", 1.0, 2.0, 3.0, time.Now())
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Equal(t, "invalid MAC address", err.Error())
}
