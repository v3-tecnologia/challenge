package repository_test

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/infrastructure/repository"
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupPicturesTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err, "Failed to create mock DB")

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err, "Failed to open GORM connection")

	return db, mock, func() {
		mockDB.Close()
	}
}

func TestPicturesRepository_CreatePictures_Success(t *testing.T) {
	db, mock, cleanup := setupPicturesTestDB(t)
	defer cleanup()

	repo := repository.NewPicturesRepository(db)
	ctx := context.Background()

	picture := &entity.Picture{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		PictureURL:       "https://example.com/image.jpg",
		PictureFormat:    "",
		RecognizedFace:   true,
		RekognitionScore: 0,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "pictures"`)).
		WithArgs(
			picture.DeviceID,
			sqlmock.AnyArg(), // CreatedAt
			sqlmock.AnyArg(), // ReceivedAt
			picture.PictureURL,
			picture.PictureFormat,
			picture.RecognizedFace,
			picture.RekognitionScore,
			picture.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(picture.ID))
	mock.ExpectCommit()

	_, err := repo.CreatePictures(ctx, picture)

	assert.NoError(t, err)
	assert.NotNil(t, picture)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestPicturesRepository_CreatePictures_Error(t *testing.T) {
	db, mock, cleanup := setupPicturesTestDB(t)
	defer cleanup()

	repo := repository.NewPicturesRepository(db)
	ctx := context.Background()

	picture := &entity.Picture{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		PictureURL:       "https://example.com/image.jpg",
		PictureFormat:    "",
		RecognizedFace:   true,
		RekognitionScore: 0,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "pictures"`)).
		WithArgs(
			picture.DeviceID,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			picture.PictureURL,
			picture.PictureFormat,
			picture.RecognizedFace,
			picture.RekognitionScore,
			picture.ID,
		).
		WillReturnError(sqlmock.ErrCancelled)
	mock.ExpectRollback()

	_, err := repo.CreatePictures(ctx, picture)

	assert.Error(t, err)
	assert.Equal(t, sqlmock.ErrCancelled, err)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
