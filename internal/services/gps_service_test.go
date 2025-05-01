package services_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"testing"
	"time"
)

func TestAddGps_Happy(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	gps := models.GpsModel{
		Latitude:  12.34,
		Longitude: 56.78,
		MAC:       "00:11:22:33:44:55",
		Timestamp: 1617181723,
	}

	mock.ExpectExec("INSERT INTO gps").
		WithArgs(gps.Latitude, gps.Longitude, gps.MAC, time.UnixMilli(gps.Timestamp)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = services.AddGps(gps)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddGps_DatabaseError(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	gps := models.GpsModel{
		Latitude:  12.34,
		Longitude: 56.78,
		MAC:       "00:11:22:33:44:55",
		Timestamp: 1617181723,
	}

	mock.ExpectExec("INSERT INTO gps").
		WithArgs(gps.Latitude, gps.Longitude, gps.MAC, time.UnixMilli(gps.Timestamp)).
		WillReturnError(fmt.Errorf("mock db error"))

	err = services.AddGps(gps)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddGps_InvalidGpsModel(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	invalidGps := models.GpsModel{
		Latitude:  0,
		Longitude: 0,
		MAC:       "",
		Timestamp: 0,
	}
	err = services.AddGps(invalidGps)

	if err == nil {
		t.Errorf("Expected error for invalid GpsModel, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}
