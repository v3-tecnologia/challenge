package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/mkafonso/go-cloud-challenge/database/sqlc"
	"github.com/mkafonso/go-cloud-challenge/entity"
)

type PostgresGPSRepository struct {
	*db.Queries
	connectionPool *pgxpool.Pool
}

func NewPostgresGPSRepository(pool *pgxpool.Pool) *PostgresGPSRepository {
	return &PostgresGPSRepository{
		connectionPool: pool,
		Queries:        db.New(pool),
	}
}

func (r *PostgresGPSRepository) SaveGPS(ctx context.Context, gps *entity.GPS) error {
	params := db.SaveGPSParams{
		ID:        gps.ID,
		DeviceID:  gps.DeviceID,
		Latitude:  *gps.Latitude,
		Longitude: *gps.Longitude,
		Timestamp: pgtype.Timestamptz{Time: gps.Timestamp, Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: gps.CreatedAt, Valid: true},
	}

	return r.Queries.SaveGPS(ctx, params)
}
