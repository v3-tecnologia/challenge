package repository

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type GPSRepositoryInterface interface {
	SaveGPS(ctx context.Context, data *entity.GPS) error
}
