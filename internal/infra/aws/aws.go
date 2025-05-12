package aws

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWSServiceInterface define os métodos necessários para o serviço AWS
type AWSServiceInterface interface {
	UploadPhoto(deviceID string, photoBytes []byte, timestamp int64) (string, error)
	ComparePhoto(deviceID, filePath string) (bool, error)
}

type AWSService struct {
	s3Client  *s3.Client
	rekClient *rekognition.Client
	bucket    string
}

func NewAWSService(bucket string) (*AWSService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}
	return &AWSService{
		s3Client:  s3.NewFromConfig(cfg),
		rekClient: rekognition.NewFromConfig(cfg),
		bucket:    bucket,
	}, nil
}

func (a *AWSService) UploadPhoto(deviceID string, photoBytes []byte, timestamp int64) (string, error) {
	if len(photoBytes) == 0 {
		return "", fmt.Errorf("empty photo data")
	}
	if len(photoBytes) > 5*1024*1024 {
		return "", fmt.Errorf("photo size exceeds 5MB")
	}

	contentType := http.DetectContentType(photoBytes)
	var extension string
	switch contentType {
	case "image/jpeg":
		extension = ".jpg"
	case "image/png":
		extension = ".png"
	default:
		return "", fmt.Errorf("unsupported photo format: %s", contentType)
	}

	key := fmt.Sprintf("%s/%d%s", deviceID, timestamp, extension)

	_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &a.bucket,
		Key:         &key,
		Body:        bytes.NewReader(photoBytes),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload photo to S3: %w", err)
	}

	return key, nil
}

func (a *AWSService) ComparePhoto(deviceID, photoKey string) (bool, error) {
	listOutput, err := a.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &a.bucket,
		Prefix: &deviceID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to list S3 objects: %w", err)
	}
	for _, obj := range listOutput.Contents {
		if *obj.Key == photoKey {
			continue
		}
		input := &rekognition.CompareFacesInput{
			SourceImage: &types.Image{
				S3Object: &types.S3Object{
					Bucket: &a.bucket,
					Name:   &photoKey,
				},
			},
			TargetImage: &types.Image{
				S3Object: &types.S3Object{
					Bucket: &a.bucket,
					Name:   obj.Key,
				},
			},
			SimilarityThreshold: aws.Float32(90.0),
		}
		output, err := a.rekClient.CompareFaces(context.TODO(), input)
		if err != nil {
			fmt.Printf("Rekognition error for %s: %v\n", *obj.Key, err)
			continue
		}
		if len(output.FaceMatches) > 0 {
			return true, nil
		}
	}
	return false, nil
}
