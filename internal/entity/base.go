package entity

import (
	"errors"
	"time"
)

var (
	ErrIDIsRequired      = errors.New("id is required")
	ErrInvalidID         = errors.New("invalid id")
	ErrMacIsRequired     = errors.New("mac address is required")
	ErrInvalidUser       = errors.New("invalid user")
	ErrInvalidCoordinate = errors.New("invalid coordinate")
	ErrInvalidTimeStamp  = errors.New("invalid timestamp")
)

func SrtToTime(datetime string) (time.Time, error) {
	layout := "2006-01-02 15:04:05" // Corresponds to "YYYY-MM-DD HH:MM:SS"
	parsedTime, err := time.Parse(layout, datetime)
	return parsedTime, err
}
