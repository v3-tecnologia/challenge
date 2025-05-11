package photos

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
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
	inputRekog := &rekognition.SearchFacesByImageInput{
		CollectionId: aws.String("your-collection-id"),
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: aws.String("your-bucket"),
				Name:   aws.String(photo.FilePath),
			},
		},
	}
	result, err := uc.RekogClient.SearchFacesByImage(context.TODO(), inputRekog)
	if err != nil {
		return nil, errors.New("failed to process photo with AWS Rekognition")
	}

	photo.Recognized = len(result.FaceMatches) > 0

	savedPhoto, err := uc.Repo.Create(photo)
	if err != nil {
		return nil, errors.New("failed to save photo data")
	}

	return savedPhoto, nil
}
