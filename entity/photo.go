package entity

import (
	"time"

	"github.com/google/uuid"
	appError "github.com/mkafonso/go-cloud-challenge/usecase/errors"
)

type Photo struct {
	ID         uuid.UUID // Unique identifier for the photo entry
	DeviceID   string    // Unique identifier of the device (MAC address)
	FilePath   string    // Path where the image is stored
	Recognized bool      // Indicates if the photo was matched using Rekognition
	Timestamp  time.Time // Timestamp when the photo was captured
	CreatedAt  time.Time // Timestamp when the entry was saved in the system
}

func NewPhoto(deviceID, filePath string, timestamp time.Time, recognized bool) (*Photo, error) {
	if deviceID == "" {
		return nil, appError.NewErrorBadRequest(
			"device_id is required",
			"please provide a valid device identifier",
		)
	}
	if filePath == "" {
		return nil, appError.NewErrorBadRequest(
			"file_path is required",
			"please provide a valid file path",
		)
	}
	if timestamp.IsZero() {
		return nil, appError.NewErrorBadRequest(
			"timestamp is required",
			"please provide a timestamp in RFC3339 format",
		)
	}

	return &Photo{
		ID:         uuid.New(),
		DeviceID:   deviceID,
		FilePath:   filePath,
		Timestamp:  timestamp,
		Recognized: recognized,
		CreatedAt:  time.Now(),
	}, nil
}
