package model

import (
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
)

type GPS struct {
	ID         uuid.UUID
	MacAddress string
	Latitude   *float64
	Longitude  *float64
	Timestamp  time.Time
	CreatedAt  time.Time
}
type GPSBuilder struct {
	macAddress string
	latitude   *float64
	longitude  *float64
	timestamp  time.Time
	err        error
}

func NewGPSBuilder() *GPSBuilder {
	return &GPSBuilder{}
}

func (b *GPSBuilder) WithMacAddress(macAddress string) *GPSBuilder {
	if b.err != nil {
		return b
	}

	if macAddress == "" {
		b.err = errors.NewErrorBadRequest(
			"missing_device_id",
			"device identifier (MAC address) is required",
		)
		return b
	}

	b.macAddress = macAddress
	return b
}

func (b *GPSBuilder) WithCoordinates(latitude, longitude *float64) *GPSBuilder {
	if b.err != nil {
		return b
	}

	if latitude == nil {
		b.err = errors.NewErrorBadRequest(
			"missing_latitude",
			"latitude coordinate is required",
		)
		return b
	}

	if longitude == nil {
		b.err = errors.NewErrorBadRequest(
			"missing_longitude",
			"longitude coordinate is required",
		)
		return b
	}

	b.latitude = latitude
	b.longitude = longitude
	return b
}

func (b *GPSBuilder) WithTimestamp(timestamp time.Time) *GPSBuilder {
	if b.err != nil {
		return b
	}

	if timestamp.IsZero() {
		b.err = errors.NewErrorBadRequest(
			"invalid_timestamp",
			"timestamp is required and must be valid",
		)
		return b
	}

	b.timestamp = timestamp
	return b
}

func (b *GPSBuilder) Build() (*GPS, error) {
	if b.err != nil {
		return nil, b.err
	}

	return &GPS{
		ID:         uuid.New(),
		MacAddress: b.macAddress,
		Latitude:   b.latitude,
		Longitude:  b.longitude,
		Timestamp:  b.timestamp,
		CreatedAt:  time.Now(),
	}, nil
}
