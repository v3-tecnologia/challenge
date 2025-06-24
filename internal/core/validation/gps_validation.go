package validation

import (
	"time"

	"github.com/dryingcore/v3-challenge/internal/config"
	"github.com/dryingcore/v3-challenge/internal/core/errors"
	"github.com/dryingcore/v3-challenge/internal/core/model"
)

func ValidateGPS(data model.GPSData) error {
	if data.MacAddress == "" {
		return errors.ErrEmptyMacAddress
	}

	if data.Timestamp.IsZero() {
		return errors.ErrTimestampIsEmpty
	}

	if data.Timestamp.After(time.Now().Add(config.AllowedSkew)) {
		return errors.ErrTimestampInFuture
	}
	return nil
}
