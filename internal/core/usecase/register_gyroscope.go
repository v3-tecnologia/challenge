package usecase

import (
	"context"

	"github.com/dryingcore/v3-challenge/internal/core/model"
	"github.com/dryingcore/v3-challenge/internal/core/ports/queue"
	"github.com/dryingcore/v3-challenge/internal/core/validation"
)

type GyroscopeUC struct {
	Publisher queue.Publisher
}

type GyroscopeUseCase interface {
	Register(data model.Gyroscope) error
}

func NewGyroscopeUC(pub queue.Publisher) GyroscopeUC {
	return GyroscopeUC{Publisher: pub}
}

func (uc *GyroscopeUC) Register(data model.Gyroscope) error {
	if err := validation.ValidateGyroscope(data); err != nil {
		return err
	}

	msg := queue.Message{
		Type: "gyroscope",
		Data: data,
	}

	return uc.Publisher.Publish(context.Background(), "telemetry.gyroscope", msg)
}
