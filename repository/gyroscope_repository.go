package repository

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/entity"
)

type GyroscopeRepositoryInterface interface {
	SaveGyroscope(ctx context.Context, data *entity.Gyroscope) error
}
