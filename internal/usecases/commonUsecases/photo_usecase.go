package commonUsecases

import (
	"mime/multipart"
	"time"
	"v3-test/internal/enums"
	local_file_storage "v3-test/internal/infra/storage"
	"v3-test/internal/interfaces"
	"v3-test/internal/models/commonModels"
	"v3-test/internal/repositories/commonRepositories"
)

type PhotoUsecase struct {
	repo    commonRepositories.PhotoRepository
	storage interfaces.FileStorageInterface
}

func NewPhotoUsecase(repo commonRepositories.PhotoRepository) PhotoUsecase {
	return PhotoUsecase{repo: repo, storage: local_file_storage.NewLocalPhotoStorage("files/photos")}
}

func (uc *PhotoUsecase) UploadPhoto(file *multipart.FileHeader, entity enums.PhotoEntity) (commonModels.PhotoModel, error) {
	filename := time.Now().Format("20060102_150405") + "-" + file.Filename

	filePath, err := uc.storage.UploadFile(file, filename)
	if err != nil {
		return commonModels.PhotoModel{}, err
	}

	photo := commonModels.PhotoModel{
		Entity:    entity,
		Url:       filePath,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uc.repo.CreatePhoto(photo)
}
