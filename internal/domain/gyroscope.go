package domain

import (
	"errors"
)

var (
	ErrDeviceIDGyroscope      = errors.New("device ID not found")
	ErrTimestampGyroscope     = errors.New("timestamp not found")
	ErrInvalidGyroscopeValues = errors.New("invalid gyroscope values")
)
