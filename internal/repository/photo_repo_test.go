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

func TestSavePhoto(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer mockDB.Close()
	db.SetConnection(mockDB)

	photo := &model.Photo{
		ID:         uuid.New(),
		MacAddress: "AA:BB:CC:DD:EE:FF",
		FileURL:    "/photos/photo1.jpg",
		IsMatch:    true,
		Timestamp:  time.Now(),
		CreatedAt:  time.Now(),
	}

	mock.ExpectExec("INSERT INTO photo_data").
		WithArgs(
			photo.ID,
			photo.MacAddress,
			photo.FileURL,
			photo.IsMatch,
			photo.Timestamp,
			photo.CreatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.SavePhoto(photo)
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectativas não foram atendidas: %v", err)
	}
}

func TestGetPhotoByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco: %v", err)
	}
	defer mockDB.Close()
	db.SetConnection(mockDB)

	photo := &model.Photo{
		ID:         uuid.New(),
		MacAddress: "AA:BB:CC:DD:EE:FF",
		FileURL:    "/photos/photo1.jpg",
		IsMatch:    true,
		Timestamp:  time.Now(),
		CreatedAt:  time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "mac_address", "file_path", "recognized", "timestamp", "created_at"}).
		AddRow(photo.ID, photo.MacAddress, photo.FileURL, photo.IsMatch, photo.Timestamp, photo.CreatedAt)

	mock.ExpectQuery("SELECT id, mac_address, file_path, recognized, timestamp, created_at FROM photo_data WHERE id =").
		WithArgs(photo.ID).
		WillReturnRows(rows)

	result, err := repository.GetPhotoByID(photo.ID)
	if err != nil {
		t.Errorf("esperava erro nil, mas recebeu: %v", err)
	}

	if result == nil {
		t.Error("esperava resultado, mas recebeu nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectativas não foram atendidas: %v", err)
	}
}
