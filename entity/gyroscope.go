package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Gyroscope struct {
	ID        uuid.UUID // Unique identifier for the gyroscope entry
	DeviceID  string    // Unique identifier of the device (MAC address)
	X         float64   // Gyroscope reading on the X axis
	Y         float64   // Gyroscope reading on the Y axis
	Z         float64   // Gyroscope reading on the Z axis
	Timestamp time.Time // Timestamp when the data was collected
	CreatedAt time.Time // Timestamp when the entry was saved in the system
}

func NewGyroscope(deviceID string, x, y, z float64, timestamp time.Time) (*Gyroscope, error) {
	if deviceID == "" {
		return nil, errors.New("deviceID cannot be empty")
	}
	if timestamp.IsZero() {
		return nil, errors.New("timestamp is required")
	}

	return &Gyroscope{
		ID:        uuid.New(),
		DeviceID:  deviceID,
		X:         x,
		Y:         y,
		Z:         z,
		Timestamp: timestamp,
		CreatedAt: time.Now(),
	}, nil
}
