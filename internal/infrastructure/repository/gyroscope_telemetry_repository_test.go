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

func setupGyroscopeTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
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

func TestGyroscopeRepository_CreateGyroscopeTelemetry_Success(t *testing.T) {
	db, mock, cleanup := setupGyroscopeTestDB(t)
	defer cleanup()

	repo := repository.NewGyroscopeTelemetryRepository(db)
	ctx := context.Background()

	gyroscope := &entity.GyroscopeTelemetry{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		X: 1,
		Y: 2,
		Z: 3,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "gyroscope_telemetry"`)).
		WithArgs(
			gyroscope.DeviceID,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			gyroscope.X,
			gyroscope.Y,
			gyroscope.Z,
			gyroscope.ID,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(gyroscope.ID))
	mock.ExpectCommit()

	err := repo.CreateGyroscopeTelemetry(ctx, gyroscope)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestGyroscopeRepository_CreateGyroscopeTelemetry_Error(t *testing.T) {
	db, mock, cleanup := setupGyroscopeTestDB(t)
	defer cleanup()

	repo := repository.NewGyroscopeTelemetryRepository(db)
	ctx := context.Background()

	gyroscope := &entity.GyroscopeTelemetry{
		BaseEntity: entity.BaseEntity{
			ID:         uuid.New(),
			DeviceID:   "00:11:22:33:44:55",
			CreatedAt:  time.Now(),
			ReceivedAt: time.Now(),
		},
		X: 1,
		Y: 2,
		Z: 3,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "gyroscope_telemetry"`)).
		WithArgs(
			gyroscope.DeviceID,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			gyroscope.X,
			gyroscope.Y,
			gyroscope.Z,
			gyroscope.ID,
		).
		WillReturnError(sqlmock.ErrCancelled)
	mock.ExpectRollback()

	err := repo.CreateGyroscopeTelemetry(ctx, gyroscope)

	assert.Error(t, err)
	assert.Equal(t, sqlmock.ErrCancelled, err)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
