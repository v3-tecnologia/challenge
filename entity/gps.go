package entity

import (
	"time"

	"github.com/google/uuid"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

type GPS struct {
	ID        uuid.UUID // Unique identifier for the GPS entry
	DeviceID  string    // Unique identifier of the device (MAC address)
	Latitude  *float64  // Latitude coordinate
	Longitude *float64  // Longitude coordinate
	Timestamp time.Time // Timestamp when the data was collected
	CreatedAt time.Time // Timestamp when the entry was saved in the system
}

func NewGPS(deviceID string, latitude, longitude *float64, timestamp time.Time) (*GPS, error) {
	if deviceID == "" {
		return nil, appError.NewErrorBadRequest(
			"device_id is required",
			"please provide a valid device identifier",
		)
	}

	if latitude == nil {
		return nil, appError.NewErrorBadRequest(
			"latitude is required",
			"please provide a valid latitude",
		)
	}

	if longitude == nil {
		return nil, appError.NewErrorBadRequest(
			"longitude is required",
			"please provide a valid longitude",
		)
	}

	if timestamp.IsZero() {
		return nil, appError.NewErrorBadRequest(
			"timestamp is required",
			"please provide a timestamp in RFC3339 format",
		)
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
