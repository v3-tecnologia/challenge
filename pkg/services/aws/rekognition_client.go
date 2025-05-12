// internal/infrastructure/aws/rekognition_client.go

package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"strings"
)

type RekognitionClient interface {
	CompareFaces(ctx context.Context, s3Key string, deviceID string) (bool, float64, error)
	IndexFace(ctx context.Context, s3Key string, deviceID string) (string, error)
}

type RekognitionClientImpl struct {
	s3Client     *s3.Client
	client       *rekognition.Client
	bucket       string
	collectionId string
}

func NewRekognitionClient() (*RekognitionClientImpl, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, err
	}

	client := rekognition.NewFromConfig(awsConfig)

	return &RekognitionClientImpl{
		client:       client,
		bucket:       "telemetry-photos-challenge",
		collectionId: "driver-faces",
		s3Client:     s3.NewFromConfig(awsConfig),
	}, nil

}

func (r *RekognitionClientImpl) CompareFaces(ctx context.Context, sourceS3Key string, deviceId string) (bool, float64, error) {
	sourceImage := &types.Image{
		S3Object: &types.S3Object{
			Bucket: aws.String(r.bucket),
			Name:   aws.String(sourceS3Key),
		},
	}

	recentImages, err := r.getRecentTrainingImages(ctx, deviceId, sourceS3Key)
	if err != nil {
		return false, 0.0, err
	}

	if len(recentImages) == 0 {
		log.Println("Nenhuma imagem encontrada")
		return false, 0.0, nil
	}

	const similarityThreshold float32 = 80.0
	var highestSimilarity float32 = 0

	for _, targetS3Key := range recentImages {
		targetImage := &types.Image{
			S3Object: &types.S3Object{
				Bucket: aws.String(r.bucket),
				Name:   aws.String(targetS3Key),
			},
		}

		input := &rekognition.CompareFacesInput{
			SourceImage:         sourceImage,
			TargetImage:         targetImage,
			SimilarityThreshold: aws.Float32(similarityThreshold),
		}

		result, err := r.client.CompareFaces(ctx, input)
		if err != nil {
			continue
		}

		for _, match := range result.FaceMatches {
			if *match.Similarity > highestSimilarity {
				highestSimilarity = *match.Similarity
			}
		}
	}

	return highestSimilarity >= similarityThreshold, float64(highestSimilarity), nil
}

func (c *RekognitionClientImpl) IndexFace(ctx context.Context, s3Key string, deviceID string) (string, error) {
	collectionName := "driver-faces"
	_, err := c.client.DescribeCollection(ctx, &rekognition.DescribeCollectionInput{
		CollectionId: aws.String(collectionName),
	})

	if err != nil {
		var resourceNotFound *types.ResourceNotFoundException
		if errors.As(err, &resourceNotFound) {
			_, err = c.client.CreateCollection(ctx, &rekognition.CreateCollectionInput{
				CollectionId: aws.String(collectionName),
			})

			if err != nil {
				return "", fmt.Errorf("falha ao criar coleção: %w", err)
			}
		} else {
			return "", fmt.Errorf("erro ao verificar coleção: %w", err)
		}
	}

	input := &rekognition.IndexFacesInput{
		CollectionId: aws.String(collectionName),
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: aws.String(c.bucket),
				Name:   aws.String(s3Key),
			},
		},
		QualityFilter: types.QualityFilterAuto,
	}

	result, err := c.client.IndexFaces(ctx, input)
	if err != nil {
		return "", err
	}

	if len(result.FaceRecords) == 0 {
		return "", fmt.Errorf("nenhuma face detectada na imagem")
	}

	return *result.FaceRecords[0].Face.FaceId, nil
}

func (r *RekognitionClientImpl) getRecentTrainingImages(ctx context.Context, deviceId string, sourceS3Key string) ([]string, error) {
	prefix := "photos/" + deviceId + "/"

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(r.bucket),
		Prefix: aws.String(prefix),
	}

	result, err := r.s3Client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar objetos no S3: %w", err)
	}

	var imageKeys []string

	for _, object := range result.Contents {
		if *object.Key != sourceS3Key {
			key := *object.Key
			if strings.HasSuffix(strings.ToLower(key), ".jpg") ||
				strings.HasSuffix(strings.ToLower(key), ".jpeg") ||
				strings.HasSuffix(strings.ToLower(key), ".png") {
				imageKeys = append(imageKeys, key)
			}
		}
	}

	return imageKeys, nil
}
