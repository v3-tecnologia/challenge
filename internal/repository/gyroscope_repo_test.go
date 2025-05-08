package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/bielgennaro/v3-challenge-cloud/internal/repository"
	"github.com/google/uuid"

	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
)

func TestSaveGyroscope(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer mockDB.Close()

	db.SetConnection(mockDB)

	gyro := &model.Gyroscope{
		ID:         uuid.New(),
		MacAddress: "00:11:22:33:44:55",
		AxisX:      1.23,
		AxisY:      4.56,
		AxisZ:      7.89,
		Timestamp:  time.Now(),
		CreatedAt:  time.Now(),
	}

	mock.ExpectExec("INSERT INTO gyroscope_data").
		WithArgs(
			gyro.ID,
			gyro.MacAddress,
			gyro.AxisX,
			gyro.AxisY,
			gyro.AxisZ,
			gyro.Timestamp,
			gyro.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.SaveGyroscope(gyro)
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("nem todas as expectativas foram atendidas: %v", err)
	}
}
