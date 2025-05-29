package services_test

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KaiRibeiro/challenge/internal/models"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockUploader struct {
	Fail bool
}

func (m *MockUploader) PutPhoto(ctx context.Context, filename string, image []byte) (string, error) {
	if m.Fail {
		return "", errors.New("mock S3 upload failure")
	}
	return "https://mockbucket.s3.region.amazonaws.com/" + filename, nil
}

type MockComparer struct {
	Recognized bool
	Err        error
}

func (m *MockComparer) Compare(ctx context.Context, mac string, filename string) (bool, error) {
	return m.Recognized, m.Err
}

func TestAddPhoto_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "AA:BB:CC:DD", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{},
		FaceComparer: &MockComparer{Recognized: true},
	}

	img := base64.StdEncoding.EncodeToString([]byte("testimage"))
	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: img,
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.NoError(t, err)
	assert.True(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddPhoto_Base64DecodeFailure(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{},
		FaceComparer: &MockComparer{},
	}

	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: "!!invalid_base64!!",
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.Error(t, err)
	assert.False(t, result)
}

func TestAddPhoto_S3UploadFailure(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{Fail: true},
		FaceComparer: &MockComparer{},
	}

	img := base64.StdEncoding.EncodeToString([]byte("testimage"))
	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: img,
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.Error(t, err)
	assert.False(t, result)
}

func TestAddPhoto_DBInsertFailure(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "AA:BB:CC:DD", sqlmock.AnyArg()).
		WillReturnError(errors.New("mock DB error"))

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{},
		FaceComparer: &MockComparer{},
	}

	img := base64.StdEncoding.EncodeToString([]byte("testimage"))
	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: img,
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.Error(t, err)
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddPhoto_RekognitionFailure(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "AA:BB:CC:DD", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{},
		FaceComparer: &MockComparer{Err: errors.New("rekognition fail")},
	}

	img := base64.StdEncoding.EncodeToString([]byte("testimage"))
	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: img,
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.Error(t, err)
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddPhoto_NotRecognized(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec("INSERT INTO photo").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "AA:BB:CC:DD", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	service := &services.PhotoDBService{
		DB:           db,
		Uploader:     &MockUploader{},
		FaceComparer: &MockComparer{Recognized: false},
	}

	img := base64.StdEncoding.EncodeToString([]byte("testimage"))
	photo := models.PhotoModel{
		MAC:         "AA:BB:CC:DD",
		ImageBase64: img,
		Timestamp:   time.Now().UnixMilli(),
	}

	result, err := service.AddPhoto(photo)
	assert.NoError(t, err)
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}
