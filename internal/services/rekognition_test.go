package services

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	rt "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/ricardoraposo/challenge/internal/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRekognitionClient struct {
	mock.Mock
}

func (m *MockRekognitionClient) DescribeCollection(ctx context.Context, params *rekognition.DescribeCollectionInput, optFns ...func(*rekognition.Options)) (*rekognition.DescribeCollectionOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.DescribeCollectionOutput), args.Error(1)
}

func (m *MockRekognitionClient) CreateCollection(ctx context.Context, params *rekognition.CreateCollectionInput, optFns ...func(*rekognition.Options)) (*rekognition.CreateCollectionOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.CreateCollectionOutput), args.Error(1)
}

func (m *MockRekognitionClient) IndexFaces(ctx context.Context, params *rekognition.IndexFacesInput, optFns ...func(*rekognition.Options)) (*rekognition.IndexFacesOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.IndexFacesOutput), args.Error(1)
}

func (m *MockRekognitionClient) SearchFacesByImage(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
	args := m.Called(ctx, params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*rekognition.SearchFacesByImageOutput), args.Error(1)
}

type TestRekognitionClient struct {
	client *MockRekognitionClient
}

func (rc *TestRekognitionClient) HandleFaceRecognition(ctx context.Context, imageKey string) (*[]interfaces.FaceMatch, error) {
	if _, err := rc.createCollection(); err != nil {
		return &[]interfaces.FaceMatch{}, err
	}

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

func (rc *TestRekognitionClient) indexFaces(ctx context.Context, imageKey string) error {
	_, err := rc.client.IndexFaces(ctx, &rekognition.IndexFacesInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		Image: &rt.Image{
			S3Object: &rt.S3Object{
				Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
				Name:   aws.String(imageKey),
			},
		},
		DetectionAttributes: []rt.Attribute{"ALL"},
		ExternalImageId:     aws.String(imageKey),
	})
	return err
}

func (rc *TestRekognitionClient) detectFace(ctx context.Context, imageKey string) (*[]interfaces.FaceMatch, error) {
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

func (rc *TestRekognitionClient) createCollection() (string, error) {
	desc, err := rc.client.DescribeCollection(context.Background(), &rekognition.DescribeCollectionInput{
		CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
	})
	if err == nil {
		return *desc.CollectionARN, nil
	}

	var notFoundErr *rt.ResourceNotFoundException
	if errors.As(err, &notFoundErr) {
		col, err := rc.client.CreateCollection(context.Background(), &rekognition.CreateCollectionInput{
			CollectionId: aws.String(os.Getenv("AWS_FACE_COLLECTION_ID")),
		})
		if err != nil {
			return "", err
		}

		return *col.CollectionArn, nil
	}

	return "", err
}

func TestMain(m *testing.M) {
	os.Setenv("AWS_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_BUCKET_KEY_PREFIX", "prefix/")
	os.Setenv("AWS_FACE_COLLECTION_ID", "test-collection")
	defer func() {
		os.Unsetenv("AWS_BUCKET_NAME")
		os.Unsetenv("AWS_BUCKET_KEY_PREFIX")
		os.Unsetenv("AWS_FACE_COLLECTION_ID")
	}()

	os.Exit(m.Run())
}

func TestHandleFaceRecognition_Success(t *testing.T) {
	t.Parallel()

	mockClient := new(MockRekognitionClient)
	rekClient := &TestRekognitionClient{
		client: mockClient,
	}

	ctx := context.Background()
	testImageKey := "test-image.png"
	collectionArn := "arn:aws:rekognition:us-east-1:123456789012:collection/test-collection"
	faceId := "face-id-123"
	similarity := float32(98.5)

	mockClient.On("DescribeCollection", mock.Anything, mock.Anything).Return(&rekognition.DescribeCollectionOutput{
		CollectionARN: aws.String(collectionArn),
	}, nil)

	mockClient.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(&rekognition.SearchFacesByImageOutput{
		FaceMatches: []rt.FaceMatch{
			{
				Similarity: aws.Float32(similarity),
				Face: &rt.Face{
					FaceId: aws.String(faceId),
				},
			},
		},
	}, nil)

	mockClient.On("IndexFaces", mock.Anything, mock.Anything).Return(&rekognition.IndexFacesOutput{}, nil)

	result, err := rekClient.HandleFaceRecognition(ctx, testImageKey)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 1)
	assert.Equal(t, float64(similarity), (*result)[0].Similarity)
	assert.Equal(t, faceId, (*result)[0].FaceID)

	mockClient.AssertExpectations(t)
}

func TestHandleFaceRecognition_CreateCollection(t *testing.T) {
	t.Parallel()

	mockClient := new(MockRekognitionClient)
	rekClient := &TestRekognitionClient{
		client: mockClient,
	}

	ctx := context.Background()
	testImageKey := "test-image.png"
	collectionArn := "arn:aws:rekognition:us-east-1:123456789012:collection/test-collection"
	faceId := "face-id-123"
	similarity := float32(98.5)

	notFoundErr := &rt.ResourceNotFoundException{
		Message: aws.String("Collection not found"),
	}
	mockClient.On("DescribeCollection", mock.Anything, mock.Anything).Return(nil, notFoundErr)

	mockClient.On("CreateCollection", mock.Anything, mock.Anything).Return(&rekognition.CreateCollectionOutput{
		CollectionArn: aws.String(collectionArn),
	}, nil)

	mockClient.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(&rekognition.SearchFacesByImageOutput{
		FaceMatches: []rt.FaceMatch{
			{
				Similarity: aws.Float32(similarity),
				Face: &rt.Face{
					FaceId: aws.String(faceId),
				},
			},
		},
	}, nil)

	mockClient.On("IndexFaces", mock.Anything, mock.Anything).Return(&rekognition.IndexFacesOutput{}, nil)

	result, err := rekClient.HandleFaceRecognition(ctx, testImageKey)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 1)
	assert.Equal(t, float64(similarity), (*result)[0].Similarity)
	assert.Equal(t, faceId, (*result)[0].FaceID)

	mockClient.AssertExpectations(t)
}

func TestHandleFaceRecognition_CreateCollectionError(t *testing.T) {
	t.Parallel()

	mockClient := new(MockRekognitionClient)
	rekClient := &TestRekognitionClient{
		client: mockClient,
	}

	ctx := context.Background()
	testImageKey := "test-image.png"
	expectedError := errors.New("failed to describe collection")

	mockClient.On("DescribeCollection", mock.Anything, mock.Anything).Return(nil, expectedError)

	result, err := rekClient.HandleFaceRecognition(ctx, testImageKey)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	mockClient.AssertExpectations(t)
}

func TestHandleFaceRecognition_DetectFaceError(t *testing.T) {
	t.Parallel()

	mockClient := new(MockRekognitionClient)
	rekClient := &TestRekognitionClient{
		client: mockClient,
	}

	ctx := context.Background()
	testImageKey := "test-image.png"
	collectionArn := "arn:aws:rekognition:us-east-1:123456789012:collection/test-collection"
	expectedError := errors.New("failed to search faces")

	mockClient.On("DescribeCollection", mock.Anything, mock.Anything).Return(&rekognition.DescribeCollectionOutput{
		CollectionARN: aws.String(collectionArn),
	}, nil)

	mockClient.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(nil, expectedError)

	result, err := rekClient.HandleFaceRecognition(ctx, testImageKey)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	mockClient.AssertExpectations(t)
}

func TestHandleFaceRecognition_IndexFacesError(t *testing.T) {
	t.Parallel()

	mockClient := new(MockRekognitionClient)
	rekClient := &TestRekognitionClient{
		client: mockClient,
	}

	ctx := context.Background()
	testImageKey := "test-image.png"
	collectionArn := "arn:aws:rekognition:us-east-1:123456789012:collection/test-collection"
	faceId := "face-id-123"
	similarity := float32(98.5)
	expectedError := errors.New("failed to index faces")

	mockClient.On("DescribeCollection", mock.Anything, mock.Anything).Return(&rekognition.DescribeCollectionOutput{
		CollectionARN: aws.String(collectionArn),
	}, nil)

	mockClient.On("SearchFacesByImage", mock.Anything, mock.Anything).Return(&rekognition.SearchFacesByImageOutput{
		FaceMatches: []rt.FaceMatch{
			{
				Similarity: aws.Float32(similarity),
				Face: &rt.Face{
					FaceId: aws.String(faceId),
				},
			},
		},
	}, nil)

	mockClient.On("IndexFaces", mock.Anything, mock.Anything).Return(nil, expectedError)

	result, err := rekClient.HandleFaceRecognition(ctx, testImageKey)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	mockClient.AssertExpectations(t)
}
