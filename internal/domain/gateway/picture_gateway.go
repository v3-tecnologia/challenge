package gateway

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
)

type PictureGateway interface {
	CreatePictures(ctx context.Context, entity *entity.Picture) (*entity.Picture, error)
}
