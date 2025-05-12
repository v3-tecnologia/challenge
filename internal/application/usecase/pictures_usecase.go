package usecase

import (
	"challenge-v3-backend/internal/domain/entity"
	"challenge-v3-backend/internal/domain/gateway"
	"challenge-v3-backend/internal/interface/dto"
	"challenge-v3-backend/pkg/services/aws"
	"context"
)

type PicturesUseCase interface {
	Create(ctx context.Context, input dto.CreatePictureRequestDTO) (dto.CreatePictureResponseDTO, error)
}

type PicturesUseCaseImpl struct {
	gateway     gateway.PictureGateway
	s3Client    services.S3Client
	rekognition services.RekognitionClient
}

func NewPicturesUseCase(gateway gateway.PictureGateway, s3Client services.S3Client, rekognition services.RekognitionClient) *PicturesUseCaseImpl {
	return &PicturesUseCaseImpl{
		gateway:     gateway,
		s3Client:    s3Client,
		rekognition: rekognition,
	}
}

func (uc *PicturesUseCaseImpl) Create(ctx context.Context, input dto.CreatePictureRequestDTO) (dto.CreatePictureResponseDTO, error) {
	s3Key, err := uc.s3Client.UploadImage(ctx, input.PictureData, input.DeviceId)

	if err != nil {
		return dto.CreatePictureResponseDTO{}, err
	}

	pictureUrl, err := uc.s3Client.GetSignedURL(ctx, s3Key)

	if err != nil {
		return dto.CreatePictureResponseDTO{}, err
	}

	_, err = uc.rekognition.IndexFace(ctx, s3Key, input.DeviceId)

	recognizedFace := false

	var rekognitionScore = 0.0

	if err != nil {
		return dto.CreatePictureResponseDTO{}, err
	}

	faceIsRecognizedFace, faceRecognizedPercent, err := uc.rekognition.CompareFaces(ctx, s3Key, input.DeviceId)

	if err != nil {
		return dto.CreatePictureResponseDTO{}, err
	}

	if faceIsRecognizedFace {
		recognizedFace = faceIsRecognizedFace
		rekognitionScore = faceRecognizedPercent
	}

	buildEntity := entity.BuildPictures(input.DeviceId, input.CreatedAt, pictureUrl, input.PictureType, recognizedFace, rekognitionScore)

	pictureCreated, err := uc.gateway.CreatePictures(ctx, buildEntity)

	if err != nil {
		return dto.CreatePictureResponseDTO{}, err
	}

	result := dto.CreatePictureResponseDTO{
		Id:           pictureCreated.ID.String(),
		DeviceId:     pictureCreated.DeviceID,
		CreatedAt:    pictureCreated.CreatedAt,
		ReceivedAt:   pictureCreated.ReceivedAt,
		IsRecognized: pictureCreated.RecognizedFace,
		PictureURL:   pictureCreated.PictureURL,
	}

	return result, nil
}
