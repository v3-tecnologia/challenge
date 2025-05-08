package gyroscope

import (
	"errors"
	"time"

	"github.com/iamrosada0/v3/internal/adapter/uuid"
	"github.com/iamrosada0/v3/internal/domain/device"
)

var (
	ErrDeviceIDGyroscope  = errors.New("DeviceID not found")
	ErrTimestampGyroscope = errors.New("timestamp not found")
)

type GyroscopeData struct {
	ID        string
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

func NewGyroscopeData(d *GyroscopeDto) (*GyroscopeData, error) {
	id := uuid.NewAdapter().Generate()

	dev, err := device.NewDevice(d.DeviceID)
	if err != nil {
		return nil, err
	}

	timestamp := time.Unix(d.Timestamp, 0)
	if timestamp.IsZero() {
		return nil, ErrTimestampGyroscope
	}

	return &GyroscopeData{
		ID:        id,
		Device:    dev,
		Timestamp: timestamp,
		X:         d.X,
		Y:         d.Y,
		Z:         d.Z,
	}, nil
}
