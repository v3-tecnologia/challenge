package config

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mkafonso/go-cloud-challenge/repository/postgres"
)

type Repositories struct {
	GPSRepo       *postgres.PostgresGPSRepository
	GyroscopeRepo *postgres.PostgresGyroscopeRepository
	PhotoRepo     *postgres.PostgresPhotoRepository
}

func SetupRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		GPSRepo:       postgres.NewPostgresGPSRepository(pool),
		GyroscopeRepo: postgres.NewPostgresGyroscopeRepository(pool),
		PhotoRepo:     postgres.NewPostgresPhotoRepository(pool),
	}
}
