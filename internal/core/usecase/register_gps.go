package usecase

import (
	"context"

	"github.com/dryingcore/v3-challenge/internal/core/model"
	"github.com/dryingcore/v3-challenge/internal/core/ports/queue"
	"github.com/dryingcore/v3-challenge/internal/core/validation"
)

type GPSUC struct {
	Publisher queue.Publisher
}

func NewGPSUseCase(pub queue.Publisher) GPSUC {
	return GPSUC{Publisher: pub}
}

func (uc *GPSUC) Register(data model.GPSData) error {
	if err := validation.ValidateGPS(data); err != nil {
		return err
	}

	msg := queue.Message{
		Type: "gps",
		Data: data,
	}

	return uc.Publisher.Publish(context.Background(), "telemetry.gps", msg)
}
