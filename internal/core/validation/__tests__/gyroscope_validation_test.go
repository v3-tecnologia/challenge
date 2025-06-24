package tests

import (
	"testing"
	"time"

	"github.com/dryingcore/v3-challenge/internal/config"
	"github.com/dryingcore/v3-challenge/internal/core/errors"
	"github.com/dryingcore/v3-challenge/internal/core/model"
	"github.com/dryingcore/v3-challenge/internal/core/validation"
)

func init() {
	config.AllowedSkew = 10 * time.Second
}

func TestValidateGyroscope(t *testing.T) {
	t.Run("should return error if mac address is empty", func(t *testing.T) {
		data := model.Gyroscope{
			Device: model.Device{
				MacAddress: "",
				Timestamp:  time.Now(),
			},
			X: 0, Y: 0, Z: 0,
		}

		err := validation.ValidateGyroscope(data)
		if err != errors.ErrEmptyMacAddress {
			t.Errorf("expected ErrEmptyMacAddress, got %v", err)
		}
	})

	t.Run("should return error if timestamp is too far in the future", func(t *testing.T) {
		data := model.Gyroscope{
			Device: model.Device{
				MacAddress: "AA:BB:CC:DD:EE:FF",
				Timestamp:  time.Now().Add(11 * time.Second),
			},
			X: 0, Y: 0, Z: 0,
		}

		err := validation.ValidateGyroscope(data)
		if err != errors.ErrTimestampInFuture {
			t.Errorf("expected ErrTimestampInFuture, got %v", err)
		}
	})

	t.Run("should accept timestamp within allowed skew", func(t *testing.T) {
		data := model.Gyroscope{
			Device: model.Device{
				MacAddress: "AA:BB:CC:DD:EE:FF",
				Timestamp:  time.Now().Add(5 * time.Second),
			},
			X: 0, Y: 0, Z: 0,
		}

		err := validation.ValidateGyroscope(data)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("should accept timestamp in the past", func(t *testing.T) {
		data := model.Gyroscope{
			Device: model.Device{
				MacAddress: "AA:BB:CC:DD:EE:FF",
				Timestamp:  time.Now().Add(-30 * time.Second),
			},
			X: 0, Y: 0, Z: 0,
		}

		err := validation.ValidateGyroscope(data)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}
