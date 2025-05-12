package usecase

import (
	"fmt"
	"v3/internal/domain"
	"v3/internal/infra/aws"
	"v3/internal/repository/photo"
)

type CreatePhotoUseCase struct {
	Repo       photo.PhotoRepository
	AWSService *aws.AWSService
}

func NewCreatePhotoUseCase(repo photo.PhotoRepository, awsService *aws.AWSService) *CreatePhotoUseCase {
	return &CreatePhotoUseCase{
		Repo:       repo,
		AWSService: awsService,
	}
}

func (uc *CreatePhotoUseCase) Execute(input domain.PhotoDto, photoBytes []byte) (*domain.Photo, error) {
	photo, err := domain.NewPhotoData(&input)
	if err != nil {
		return nil, err
	}
	if len(photoBytes) == 0 {
		return nil, domain.ErrPhotoData
	}

	photoPath, err := uc.AWSService.UploadPhoto(input.DeviceID, photoBytes, input.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrProcessPhotoWithAWSRekognition, err)
	}
	photo.FilePath = photoPath

	recognized, err := uc.AWSService.ComparePhoto(input.DeviceID, photoPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrProcessPhotoWithAWSRekognition, err)
	}
	photo.Recognized = recognized

	savedPhoto, err := uc.Repo.Create(photo)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrSavePhotoData, err)
	}

	return savedPhoto, nil
}

// package usecase

// import (
// 	"fmt"
// 	"v3/internal/domain"
// 	"v3/internal/infra/aws"
// 	"v3/internal/repository/photo"
// )

// type CreatePhotoUseCase struct {
// 	Repo       photo.PhotoRepository
// 	AWSService *aws.AWSService
// }

// func NewCreatePhotoUseCase(repo photo.PhotoRepository, awsService *aws.AWSService) *CreatePhotoUseCase {
// 	return &CreatePhotoUseCase{
// 		Repo:       repo,
// 		AWSService: awsService,
// 	}
// }

// func (uc *CreatePhotoUseCase) Execute(input domain.PhotoDto) (*domain.Photo, error) {
// 	photo, err := domain.NewPhotoData(&input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Upload photo to S3
// 	photoPath, err := uc.AWSService.UploadPhoto(input.DeviceID, input.Photo, input.Timestamp)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %v", domain.ErrProcessPhotoWithAWSRekognition, err)
// 	}
// 	photo.FilePath = photoPath

// 	// Compare with previous photos
// 	recognized, err := uc.AWSService.ComparePhoto(input.DeviceID, photoPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %v", domain.ErrProcessPhotoWithAWSRekognition, err)
// 	}
// 	photo.Recognized = recognized

// 	// Save to database
// 	savedPhoto, err := uc.Repo.Create(photo)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %v", domain.ErrSavePhotoData, err)
// 	}

// 	return savedPhoto, nil
// }
