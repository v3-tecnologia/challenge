package gateway

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
)

type GyroscopeTelemetryGateway interface {
	CreateGyroscopeTelemetry(ctx context.Context, entity *entity.GyroscopeTelemetry) error
}
