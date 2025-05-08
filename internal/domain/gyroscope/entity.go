package gyroscope

import (
	"errors"
	"time"

	"github.com/iamrosada0/v3/internal/domain/device"
)

type GyroscopeData struct {
	Device    *device.Device
	X         float64
	Y         float64
	Z         float64
	Timestamp time.Time
}

func NewGyroscopeData(deviceID string, x, y, z float64, timestamp time.Time) (*GyroscopeData, error) {
	dev, err := device.NewDevice(deviceID)
	if err != nil {
		return nil, err
	}
	if timestamp.IsZero() {
		return nil, errors.New("timestamp is required")
	}
	return &GyroscopeData{Device: dev, X: x, Y: y, Z: z, Timestamp: timestamp}, nil
}
