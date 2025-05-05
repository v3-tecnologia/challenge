package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type GPS struct {
	ID        uuid.UUID // Unique identifier for the GPS entry
	DeviceID  string    // Unique identifier of the device (MAC address)
	Latitude  float64   // Latitude coordinate
	Longitude float64   // Longitude coordinate
	Timestamp time.Time // Timestamp when the data was collected
	CreatedAt time.Time // Timestamp when the entry was saved in the system
}

func NewGPS(deviceID string, latitude, longitude float64, timestamp time.Time) (*GPS, error) {
	if deviceID == "" {
		return nil, errors.New("deviceID cannot be empty")
	}
	if timestamp.IsZero() {
		return nil, errors.New("timestamp is required")
	}

	return &GPS{
		ID:        uuid.New(),
		DeviceID:  deviceID,
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: timestamp,
		CreatedAt: time.Now(),
	}, nil
}
