package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bielgennaro/v3-challenge-cloud/internal/db"
	"github.com/bielgennaro/v3-challenge-cloud/internal/model"
	"github.com/bielgennaro/v3-challenge-cloud/internal/repository"
	"github.com/google/uuid"
)

func TestSaveGPS(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer mockDB.Close()
	db.SetConnection(mockDB)

	lat := -23.55052
	lng := -46.633308

	gps := &model.GPS{
		ID:         uuid.New(),
		MacAddress: "DE:AD:BE:EF:00:01",
		Latitude:   &lat,
		Longitude:  &lng,
		Timestamp:  time.Now(),
		CreatedAt:  time.Now(),
	}

	mock.ExpectExec("INSERT INTO gps_data").
		WithArgs(
			gps.ID,
			gps.MacAddress,
			gps.Latitude,
			gps.Longitude,
			gps.Timestamp,
			gps.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.SaveGPS(gps)
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectativas não foram atendidas: %v", err)
	}
}
