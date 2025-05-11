package photos

import (
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/iamrosada0/v3/internal/domain"
	"github.com/iamrosada0/v3/internal/repository/photo"
)

type PhotoInputDto struct {
	DeviceID  string `json:"deviceId"`
	Timestamp int64  `json:"timestamp"`
	FilePath  string `json:"file_path"`
}

type CreatePhotoUseCase struct {
	Repo        photo.PhotoRepository
	RekogClient *rekognition.Client
}

func NewCreatePhotoUseCase(repo photo.PhotoRepository, rekogClient *rekognition.Client) *CreatePhotoUseCase {
	return &CreatePhotoUseCase{
		Repo:        repo,
		RekogClient: rekogClient,
	}
}

func (uc *CreatePhotoUseCase) Execute(input PhotoInputDto) (*domain.Photo, error) {

	photo, err := domain.NewPhotoData(&domain.PhotoDto{
		DeviceID:  input.DeviceID,
		Timestamp: input.Timestamp,
		FilePath:  input.FilePath,
	})
	if err != nil {
		return nil, err
	}

}
