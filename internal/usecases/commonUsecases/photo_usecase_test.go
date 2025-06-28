package commonUsecases

import (
	"errors"
	"mime/multipart"
	"testing"
	"v3-test/internal/enums"
	"v3-test/internal/models/commonModels"
)

type mockPhotoRepo struct {
	mockCreatePhoto func(commonModels.PhotoModel) (commonModels.PhotoModel, error)
}

func (m *mockPhotoRepo) CreatePhoto(photo commonModels.PhotoModel) (commonModels.PhotoModel, error) {
	return m.mockCreatePhoto(photo)
}

type mockStorage struct {
	mockUploadFile func(file *multipart.FileHeader, filename string) (string, error)
}

func (m *mockStorage) UploadFile(file *multipart.FileHeader, filename string) (string, error) {
	return m.mockUploadFile(file, filename)
}

func TestUploadPhoto_Success(t *testing.T) {
	file := &multipart.FileHeader{
		Filename: "example.jpg",
	}

	mockStorage := &mockStorage{
		mockUploadFile: func(file *multipart.FileHeader, filename string) (string, error) {
			return "files/photos/" + filename, nil
		},
	}

	mockRepo := &mockPhotoRepo{
		mockCreatePhoto: func(photo commonModels.PhotoModel) (commonModels.PhotoModel, error) {
			return photo, nil
		},
	}

	usecase := PhotoUsecase{
		repo:    mockRepo,
		storage: mockStorage,
	}

	result, err := usecase.UploadPhoto(file, enums.PhotoEntity(enums.Telemetry))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Url == "" {
		t.Error("expected non-empty URL")
	}

	if result.Entity != enums.PhotoEntity(enums.Telemetry) {
		t.Errorf("expected entity PhotoEntityProduct, got %v", result.Entity)
	}
}

func TestUploadPhoto_ErrorOnStorage(t *testing.T) {
	file := &multipart.FileHeader{
		Filename: "fail.jpg",
	}

	mockStorage := &mockStorage{
		mockUploadFile: func(file *multipart.FileHeader, filename string) (string, error) {
			return "", errors.New("upload failed")
		},
	}

	mockRepo := &mockPhotoRepo{}

	usecase := PhotoUsecase{
		repo:    mockRepo,
		storage: mockStorage,
	}

	_, err := usecase.UploadPhoto(file, enums.PhotoEntity(enums.Telemetry))

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "upload failed" {
		t.Errorf("expected 'upload failed', got '%s'", err.Error())
	}
}

func TestNewPhotoUsecase(t *testing.T) {
	mockRepo := &mockPhotoRepo{}

	usecase := NewPhotoUsecase(mockRepo)

	if usecase.repo != mockRepo {
		t.Error("expected repo to be set correctly")
	}

	if usecase.storage == nil {
		t.Fatal("expected storage to be initialized, got nil")
	}
}
