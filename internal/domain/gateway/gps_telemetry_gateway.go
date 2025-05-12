package gateway

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
)

type GPSTelemetryGateway interface {
	CreateGPSTelemetry(ctx context.Context, entity *entity.GPSTelemetry) error
}
