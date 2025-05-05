package repository

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type PhotoRepositoryInterface interface {
	SavePhoto(ctx context.Context, data *entity.Photo) error
}
