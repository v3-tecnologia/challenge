package entity

import (
	"time"

	"github.com/google/uuid"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

type Gyroscope struct {
	ID        uuid.UUID // Unique identifier for the gyroscope entry
	DeviceID  string    // Unique identifier of the device (MAC address)
	X         *float64  // Gyroscope reading on the X axis
	Y         *float64  // Gyroscope reading on the Y axis
	Z         *float64  // Gyroscope reading on the Z axis
	Timestamp time.Time // Timestamp when the data was collected
	CreatedAt time.Time // Timestamp when the entry was saved in the system
}

func NewGyroscope(deviceID string, x, y, z *float64, timestamp time.Time) (*Gyroscope, error) {
	if deviceID == "" {
		return nil, appError.NewErrorBadRequest(
			"device_id is required",
			"please provide a valid device identifier",
		)
	}

	if timestamp.IsZero() {
		return nil, appError.NewErrorBadRequest(
			"timestamp is required",
			"please provide a timestamp in RFC3339 format",
		)
	}

	if x == nil {
		return nil, appError.NewErrorBadRequest(
			"x is required",
			"please provide a valid x value",
		)
	}

	if y == nil {
		return nil, appError.NewErrorBadRequest(
			"y is required",
			"please provide a valid y value",
		)
	}

	if z == nil {
		return nil, appError.NewErrorBadRequest(
			"z is required",
			"please provide a valid z value",
		)
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
