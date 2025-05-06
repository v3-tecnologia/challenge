package factory

import (
	"github.com/mkafonso/go-cloud-challenge/recognition"
	"github.com/mkafonso/go-cloud-challenge/repository"
	"github.com/mkafonso/go-cloud-challenge/storage"
	"github.com/mkafonso/go-cloud-challenge/usecase"
)

func SavePhotoDataFactory(
	repo repository.PhotoRepositoryInterface,
	recognizer recognition.FaceRecognitionService,
	storage storage.PhotoStorageService,
) *usecase.SavePhoto {
	usecase := usecase.NewSavePhoto(repo, recognizer, storage)
	return usecase
}
