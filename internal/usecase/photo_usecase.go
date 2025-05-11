package usecase

// import (
// 	"context"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/rekognition"
// 	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
// 	"v3/internal/domain"
// 	"v3/internal/repository/photo"
// )

// type CreatePhotoUseCase struct {
// 	Repo        photo.PhotoRepository
// 	RekogClient *rekognition.Client
// }

// func NewCreatePhotoUseCase(repo photo.PhotoRepository, rekogClient *rekognition.Client) *CreatePhotoUseCase {
// 	return &CreatePhotoUseCase{
// 		Repo:        repo,
// 		RekogClient: rekogClient,
// 	}
// }

// func (uc *CreatePhotoUseCase) Execute(input domain.PhotoDto) (*domain.Photo, error) {

// 	photo, err := domain.NewPhotoData(&domain.PhotoDto{
// 		DeviceID:  input.DeviceID,
// 		Timestamp: input.Timestamp,
// 		FilePath:  input.FilePath,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	inputRekog := &rekognition.SearchFacesByImageInput{
// 		CollectionId: aws.String("your-collection-id"),
// 		Image: &types.Image{
// 			S3Object: &types.S3Object{
// 				Bucket: aws.String("your-bucket"),
// 				Name:   aws.String(photo.FilePath),
// 			},
// 		},
// 	}
// 	result, err := uc.RekogClient.SearchFacesByImage(context.TODO(), inputRekog)
// 	if err != nil {
// 		return nil, domain.ErrProcessPhotoWithAWSRekognition
// 	}

// 	photo.Recognized = len(result.FaceMatches) > 0

// 	savedPhoto, err := uc.Repo.Create(photo)
// 	if err != nil {
// 		return nil, domain.ErrSavePhotoData
// 	}

// 	return savedPhoto, nil
// }
