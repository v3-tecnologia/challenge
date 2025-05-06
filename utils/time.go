package utils

import (
	"errors"
	"time"
)

func ParseRFC3339(timestamp string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return time.Time{}, errors.New("invalid timestamp format")
	}

	return t, nil
}
