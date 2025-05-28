package services_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestAddGps_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	service := services.NewGPSDBService(db)

	gps := models.GpsModel{
		Latitude:  10.1234,
		Longitude: 20.5678,
		MAC:       "00:11:22:33:44:55",
		Timestamp: time.Now().UnixMilli(),
	}

	mock.ExpectExec("INSERT INTO gps").
		WithArgs(gps.Latitude, gps.Longitude, gps.MAC, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = service.AddGps(gps)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddGps_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	service := services.NewGPSDBService(db)

	gps := models.GpsModel{
		Latitude:  1.0,
		Longitude: 1.0,
		MAC:       "aa:bb:cc:dd:ee:ff",
		Timestamp: time.Now().UnixMilli(),
	}

	mock.ExpectExec("INSERT INTO gps").
		WithArgs(gps.Latitude, gps.Longitude, gps.MAC, sqlmock.AnyArg()).
		WillReturnError(sqlmock.ErrCancelled)

	err = service.AddGps(gps)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to insert gps data into database")
	assert.NoError(t, mock.ExpectationsWereMet())
}
