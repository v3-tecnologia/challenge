package services_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestAddGyroscope_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	service := services.NewGyroscopeDBService(db)

	gyroscope := models.GyroscopeModel{
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
		MAC:       "00:11:22:33:44:55",
		Timestamp: time.Now().UnixMilli(),
	}

	mock.ExpectExec("INSERT INTO gyroscope").
		WithArgs(gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = service.AddGyroscope(gyroscope)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddGyroscope_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	service := services.NewGyroscopeDBService(db)

	gyroscope := models.GyroscopeModel{
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
		MAC:       "00:11:22:33:44:55",
		Timestamp: time.Now().UnixMilli(),
	}

	mock.ExpectExec("INSERT INTO gyroscope").
		WithArgs(gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, sqlmock.AnyArg()).
		WillReturnError(sqlmock.ErrCancelled)

	err = service.AddGyroscope(gyroscope)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to insert gyroscope data into database")
	assert.NoError(t, mock.ExpectationsWereMet())
}
