package aws

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSService struct {
	s3Client  *s3.Client
	rekClient *rekognition.Client
	bucket    string
}

func NewAWSService(bucket string) (*AWSService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-2"), // Adjust region as needed
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
	key := fmt.Sprintf("%s/%d.jpg", deviceID, timestamp)
	_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &a.bucket,
		Key:         &key,
		Body:        bytes.NewReader(photoBytes),
		ContentType: aws.String("image/jpeg"),
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

// package aws

// import (
// 	"bytes"
// 	"context"
// 	"encoding/base64"
// 	"errors"
// 	"fmt"
// 	"strings"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/rekognition"
// 	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/aws/smithy-go"
// )

// type AWSService struct {
// 	s3Client  *s3.Client
// 	rekClient *rekognition.Client
// 	bucket    string
// }

// func NewAWSService(bucket string) (*AWSService, error) {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("us-east-1"), // Adjust region as needed
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load AWS config: %w", err)
// 	}
// 	return &AWSService{
// 		s3Client:  s3.NewFromConfig(cfg),
// 		rekClient: rekognition.NewFromConfig(cfg),
// 		bucket:    bucket,
// 	}, nil
// }

// func (a *AWSService) UploadPhoto(deviceID string, photoBase64 string, timestamp int64) (string, error) {
// 	// Remove data URI prefix if present
// 	photoBase64 = strings.TrimPrefix(photoBase64, "data:image/jpeg;base64,")
// 	photoBase64 = strings.TrimPrefix(photoBase64, "data:image/png;base64,")

// 	photoBytes, err := base64.StdEncoding.DecodeString(photoBase64)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to decode base64 photo: %w", err)
// 	}
// 	key := fmt.Sprintf("%s/%d.jpg", deviceID, timestamp)
// 	_, err = a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
// 		Bucket:      &a.bucket,
// 		Key:         &key,
// 		Body:        bytes.NewReader(photoBytes),
// 		ContentType: aws.String("image/jpeg"),
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to upload photo to S3: %w", err)
// 	}
// 	return key, nil
// }

// func (a *AWSService) ComparePhoto(deviceID, photoKey string) (bool, error) {
// 	listOutput, err := a.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
// 		Bucket: &a.bucket,
// 		Prefix: &deviceID,
// 	})
// 	if err != nil {
// 		return false, fmt.Errorf("failed to list S3 objects: %w", err)
// 	}

// 	for _, obj := range listOutput.Contents {
// 		if *obj.Key == photoKey {
// 			continue
// 		}
// 		input := &rekognition.CompareFacesInput{
// 			SourceImage: &types.Image{
// 				S3Object: &types.S3Object{
// 					Bucket: &a.bucket,
// 					Name:   &photoKey,
// 				},
// 			},
// 			TargetImage: &types.Image{
// 				S3Object: &types.S3Object{
// 					Bucket: &a.bucket,
// 					Name:   obj.Key,
// 				},
// 			},
// 			SimilarityThreshold: aws.Float32(90.0),
// 		}
// 		output, err := a.rekClient.CompareFaces(context.TODO(), input)
// 		if err != nil {
// 			// Log specific Rekognition errors for debugging
// 			var apiErr smithy.APIError
// 			if errors.As(err, &apiErr) {
// 				fmt.Printf("Rekognition error: %s - %s\n", apiErr.ErrorCode(), apiErr.ErrorMessage())
// 			}
// 			continue // Skip if comparison fails for one image
// 		}
// 		if len(output.FaceMatches) > 0 {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }
