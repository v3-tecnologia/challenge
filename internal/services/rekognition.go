package services

import (
	"context"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	rt "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/google/uuid"
	"github.com/ricardoraposo/challenge/internal/interfaces"
)

type RekognitionClient struct {
	client *rekognition.Client
}

func NewRekognitionClient() interfaces.FaceDetector {
	client := rekognition.NewFromConfig(awsConfig)

	return &RekognitionClient{
		client: client,
	}
}

func (rc *RekognitionClient) HandleFaceRecognition(
	ctx context.Context,
	imageKey string,
) (*[]interfaces.FaceMatch, error) {
	key := os.Getenv("AWS_BUCKET_KEY_PREFIX") + imageKey
	result, err := rc.detectFace(ctx, key)
	if err != nil {
		return &[]interfaces.FaceMatch{}, err
	}

	err = rc.indexFaces(ctx, key)
	if err != nil {
		return &[]interfaces.FaceMatch{}, err
	}

	return result, nil
}

func (rc *RekognitionClient) indexFaces(ctx context.Context, imageKey string) error {
	fileKey := path.Base(imageKey)
	_, err := rc.client.IndexFaces(ctx, &rekognition.IndexFacesInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		Image: &rt.Image{
			S3Object: &rt.S3Object{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
				Name:   aws.String(imageKey),
			},
		},
		DetectionAttributes: []rt.Attribute{"ALL"},
		ExternalImageId:     aws.String(fileKey),
	})

	return err
}

func (rc RekognitionClient) detectFace(ctx context.Context, imageKey string) (*[]interfaces.FaceMatch, error) {
	result, err := rc.client.SearchFacesByImage(ctx, &rekognition.SearchFacesByImageInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		Image: &rt.Image{
			S3Object: &rt.S3Object{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
				Name:   aws.String(imageKey),
			},
		},
	})

	if err != nil {
		return &[]interfaces.FaceMatch{}, err
	}

	var faceMatches []interfaces.FaceMatch
	for _, face := range result.FaceMatches {
		faceMatches = append(faceMatches, interfaces.FaceMatch{
			Similarity: float64(*face.Similarity),
			FaceID:     *face.Face.FaceId,
		})
	}

	return &faceMatches, nil
}

func (rc *RekognitionClient) CreateCollection() (string, error) {
    col, err := rc.client.CreateCollection(context.Background(), &rekognition.CreateCollectionInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
	})

    return *col.CollectionArn, err
}

func (rc *RekognitionClient) CreateUser() (string, error) {
	userId := uuid.NewString()
	_, err := rc.client.CreateUser(context.Background(), &rekognition.CreateUserInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		UserId:       aws.String(userId),
	})

	return userId, err
}

func (rc *RekognitionClient) AssociateFace(userId string, faceIds []string) (*rekognition.AssociateFacesOutput, error) {
	return rc.client.AssociateFaces(context.Background(), &rekognition.AssociateFacesInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		UserId:       aws.String(userId),
		FaceIds:      faceIds,
	})
}
