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
type GyroscopeDto struct {
	ID        string  `json:"id"`
	DeviceID  string  `json:"deviceId"`
	Timestamp int64   `json:"timestamp"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
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
