package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/mkafonso/go-cloud-challenge/database/sqlc"
	"github.com/mkafonso/go-cloud-challenge/entity"
)

type PostgresPhotoRepository struct {
	*db.Queries
	connectionPool *pgxpool.Pool
}

func NewPostgresPhotoRepository(pool *pgxpool.Pool) *PostgresPhotoRepository {
	return &PostgresPhotoRepository{
		connectionPool: pool,
		Queries:        db.New(pool),
	}
}

func (r *PostgresPhotoRepository) SavePhoto(ctx context.Context, photo *entity.Photo) error {
	params := db.SavePhotoParams{
		ID:         photo.ID,
		DeviceID:   photo.DeviceID,
		FilePath:   photo.FilePath,
		Recognized: photo.Recognized,
		Timestamp:  pgtype.Timestamptz{Time: photo.Timestamp, Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: photo.CreatedAt, Valid: true},
	}

	return r.Queries.SavePhoto(ctx, params)
}
