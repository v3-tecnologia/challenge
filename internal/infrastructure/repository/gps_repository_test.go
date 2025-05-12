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

func setupGPSTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
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

func TestGPSRepository_CreateGPSTelemetry_Success(t *testing.T) {
	// Arrange
	db, mock, cleanup := setupGPSTestDB(t)
	defer cleanup()

	repo := repository.NewGPSRepository(db)
	ctx := context.Background()

	gps := &entity.GPSTelemetry{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		Latitude:  -23.550520,
		Longitude: -46.633308,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "gps_telemetry"`)).
		WithArgs(
			gps.DeviceID,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			gps.Latitude,
			gps.Longitude,
			gps.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(gps.ID))
	mock.ExpectCommit()

	err := repo.CreateGPSTelemetry(ctx, gps)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestGPSRepository_CreateGPSTelemetry_Error(t *testing.T) {
	db, mock, cleanup := setupGPSTestDB(t)
	defer cleanup()

	repo := repository.NewGPSRepository(db)
	ctx := context.Background()

	gps := &entity.GPSTelemetry{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		Latitude:  -23.550520,
		Longitude: -46.633308,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "gps_telemetry"`)).
		WithArgs(
			gps.DeviceID,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			gps.Latitude,
			gps.Longitude,
			gps.ID,
		).
		WillReturnError(sqlmock.ErrCancelled)
	mock.ExpectRollback()

	err := repo.CreateGPSTelemetry(ctx, gps)

	assert.Error(t, err)
	assert.Equal(t, sqlmock.ErrCancelled, err)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
