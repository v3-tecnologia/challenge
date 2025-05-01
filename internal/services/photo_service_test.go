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

func TestAddPhoto_Happy(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	photo := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAABnSURBVDhP7cyhDcAgDERRZ5fMkhmYgQ1S0NLR0dHQUdHxJQcJpEBCKv7kSJZ1W3d5wQvGmBzH4TiO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4/wBZ8kQCPuQ6VQAAAAASUVORK5CYII=",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(photo.ImageBase64, photo.MAC, time.UnixMilli(photo.Timestamp)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = services.AddPhoto(photo)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddPhoto_DatabaseError(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	photo := models.PhotoModel{
		ImageBase64: "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAABnSURBVDhP7cyhDcAgDERRZ5fMkhmYgQ1S0NLR0dHQUdHxJQcJpEBCKv7kSJZ1W3d5wQvGmBzH4TiO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4ziO4/wBZ8kQCPuQ6VQAAAAASUVORK5CYII=",
		MAC:         "8C:16:45:8D:F3:7B",
		Timestamp:   1746110367207,
	}

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(photo.ImageBase64, photo.MAC, time.UnixMilli(photo.Timestamp)).
		WillReturnError(fmt.Errorf("mock db error"))

	err = services.AddPhoto(photo)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestAddPhoto_InvalidPhotoModel(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer dbCon.Close()

	db.DB = dbCon

	invalidPhoto := models.PhotoModel{
		ImageBase64: "",
		MAC:         "",
		Timestamp:   0,
	}
	err = services.AddPhoto(invalidPhoto)

	if err == nil {
		t.Errorf("Expected error for invalid PhotoModel, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}
