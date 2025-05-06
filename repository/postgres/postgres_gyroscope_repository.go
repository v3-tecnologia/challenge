package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/mkafonso/go-cloud-challenge/database/sqlc"
	"github.com/mkafonso/go-cloud-challenge/entity"
)

type PostgresGyroscopeRepository struct {
	*db.Queries
	connectionPool *pgxpool.Pool
}

func NewPostgresGyroscopeRepository(pool *pgxpool.Pool) *PostgresGyroscopeRepository {
	return &PostgresGyroscopeRepository{
		connectionPool: pool,
		Queries:        db.New(pool),
	}
}

func (r *PostgresGyroscopeRepository) SaveGyroscope(ctx context.Context, gyro *entity.Gyroscope) error {
	params := db.SaveGyroscopeParams{
		ID:        gyro.ID,
		DeviceID:  gyro.DeviceID,
		X:         *gyro.X,
		Y:         *gyro.Y,
		Z:         *gyro.Z,
		Timestamp: pgtype.Timestamptz{Time: gyro.Timestamp, Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: gyro.CreatedAt, Valid: true},
	}

	return r.Queries.SaveGyroscope(ctx, params)
}
