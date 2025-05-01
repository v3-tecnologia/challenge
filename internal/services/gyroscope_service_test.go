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

func TestAddGyroscope_Happy(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	gyroscope := models.GyroscopeModel{
		X:         12.34,
		Y:         56.78,
		Z:         321.23,
		MAC:       "00:11:22:33:44:55",
		Timestamp: 1617181723,
	}

	mock.ExpectExec("INSERT INTO gyroscope").
		WithArgs(gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, time.UnixMilli(gyroscope.Timestamp)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = services.AddGyroscope(gyroscope)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddGyroscope_DatabaseError(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	gyroscope := models.GyroscopeModel{
		X:         12.34,
		Y:         56.78,
		Z:         321.23,
		MAC:       "00:11:22:33:44:55",
		Timestamp: 1617181723,
	}

	mock.ExpectExec("INSERT INTO gyroscope").
		WithArgs(gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, time.UnixMilli(gyroscope.Timestamp)).
		WillReturnError(fmt.Errorf("mock db error"))

	err = services.AddGyroscope(gyroscope)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddGyroscope_InvalidGyroscopeModel(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	invalidGyroscope := models.GyroscopeModel{
		X:         0,
		Y:         0,
		Z:         0,
		MAC:       "",
		Timestamp: 0,
	}
	err = services.AddGyroscope(invalidGyroscope)

	if err == nil {
		t.Errorf("Expected error for invalid GyroscopeModel, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}
