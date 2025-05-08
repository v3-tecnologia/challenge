package model

import (
	"time"

	"github.com/bielgennaro/v3-challenge-cloud/internal/errors"
	"github.com/google/uuid"
)

type Gyroscope struct {
	ID         uuid.UUID
	MacAddress string
	AxisX      float64
	AxisY      float64
	AxisZ      float64
	Timestamp  time.Time
	CreatedAt  time.Time
}

type GyroscopeBuilder struct {
	macAddress string
	axisX      float64
	axisY      float64
	axisZ      float64
	timestamp  time.Time
	err        error
}

func NewGyroscopeBuilder() *GyroscopeBuilder {
	return &GyroscopeBuilder{}
}

func (b *GyroscopeBuilder) WithMacAddress(macAddress string) *GyroscopeBuilder {
	if b.err != nil {
		return b
	}

	if macAddress == "" {
		b.err = errors.NewErrorBadRequest(
			"missing_device_id",
			"device identifier is required",
		)
		return b
	}

	b.macAddress = macAddress
	return b
}

func (b *GyroscopeBuilder) WithAxisValues(x, y, z float64) *GyroscopeBuilder {
	if b.err != nil {
		return b
	}

	if x == 0 && y == 0 && z == 0 {
		b.err = errors.NewErrorBadRequest(
			"invalid_axis_values",
			"at least one axis must have a non-zero value",
		)
		return b
	}

	b.axisX = x
	b.axisY = y
	b.axisZ = z
	return b
}

func (b *GyroscopeBuilder) WithTimestamp(timestamp time.Time) *GyroscopeBuilder {
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

func (b *GyroscopeBuilder) Build() (*Gyroscope, error) {
	if b.err != nil {
		return nil, b.err
	}

	return &Gyroscope{
		ID:         uuid.New(),
		MacAddress: b.macAddress,
		AxisX:      b.axisX,
		AxisY:      b.axisY,
		AxisZ:      b.axisZ,
		Timestamp:  b.timestamp,
		CreatedAt:  time.Now(),
	}, nil
}
