package domain

import (
	"errors"
)

var (
	ErrDeviceIDGPS      = errors.New("device ID not found")
	ErrTimestampGPS     = errors.New("timestamp not found")
	ErrInvalidGPSValues = errors.New("invalid GPS values")
)
