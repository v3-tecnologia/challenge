package usecase

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"v3/internal/domain"

// 	"github.com/aws/aws-sdk-go-v2/service/rekognition"
// 	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
// )

// type mockPhotoRepository struct {
// 	createFunc func(d *domain.Photo) (*domain.Photo, error)
// }

// func (m *mockPhotoRepository) Create(d *domain.Photo) (*domain.Photo, error) {
// 	return m.createFunc(d)
// }

// type mockRekognitionClient struct {
// 	searchFacesFunc func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error)
// }

// func (m *mockRekognitionClient) SearchFacesByImage(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 	return m.searchFacesFunc(ctx, params, optFns...)
// }

// func TestCreatePhotoUseCase_Execute(t *testing.T) {
// 	validInput := PhotoInputDto{
// 		DeviceID:  "00:1A:2B:3C:4D:5E",
// 		Timestamp: time.Now().Unix(),
// 		FilePath:  "s3://bucket/photos/photo1.jpg",
// 	}

// 	tests := []struct {
// 		name           string
// 		input          PhotoInputDto
// 		repoMock       *mockPhotoRepository
// 		rekogMock      *mockRekognitionClient
// 		wantErr        bool
// 		wantRecognized bool
// 	}{
// 		{
// 			name:  "valid input, recognized",
// 			input: validInput,
// 			repoMock: &mockPhotoRepository{
// 				createFunc: func(d *domain.Photo) (*domain.Photo, error) {
// 					return d, nil
// 				},
// 			},
// 			rekogMock: &mockRekognitionClient{
// 				searchFacesFunc: func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 					return &rekognition.SearchFacesByImageOutput{
// 						FaceMatches: []types.FaceMatch{{}},
// 					}, nil
// 				},
// 			},
// 			wantErr:        false,
// 			wantRecognized: true,
// 		},
// 		{
// 			name:  "valid input, not recognized",
// 			input: validInput,
// 			repoMock: &mockPhotoRepository{
// 				createFunc: func(d *domain.Photo) (*domain.Photo, error) {
// 					return d, nil
// 				},
// 			},
// 			rekogMock: &mockRekognitionClient{
// 				searchFacesFunc: func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 					return &rekognition.SearchFacesByImageOutput{
// 						FaceMatches: []types.FaceMatch{},
// 					}, nil
// 				},
// 			},
// 			wantErr:        false,
// 			wantRecognized: false,
// 		},
// 		{
// 			name: "invalid input (invalid MAC)",
// 			input: PhotoInputDto{
// 				DeviceID:  "invalid",
// 				Timestamp: time.Now().Unix(),
// 				FilePath:  "s3://bucket/photos/photo1.jpg",
// 			},
// 			repoMock: &mockPhotoRepository{
// 				createFunc: func(d *domain.Photo) (*domain.Photo, error) {
// 					return nil, errors.New("should not be called")
// 				},
// 			},
// 			rekogMock: &mockRekognitionClient{
// 				searchFacesFunc: func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 					return nil, errors.New("should not be called")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name:  "rekognition error",
// 			input: validInput,
// 			repoMock: &mockPhotoRepository{
// 				createFunc: func(d *domain.Photo) (*domain.Photo, error) {
// 					return nil, errors.New("should not be called")
// 				},
// 			},
// 			rekogMock: &mockRekognitionClient{
// 				searchFacesFunc: func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 					return nil, errors.New("rekognition error")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name:  "repository error",
// 			input: validInput,
// 			repoMock: &mockPhotoRepository{
// 				createFunc: func(d *domain.Photo) (*domain.Photo, error) {
// 					return nil, errors.New("database error")
// 				},
// 			},
// 			rekogMock: &mockRekognitionClient{
// 				searchFacesFunc: func(ctx context.Context, params *rekognition.SearchFacesByImageInput, optFns ...func(*rekognition.Options)) (*rekognition.SearchFacesByImageOutput, error) {
// 					return &rekognition.SearchFacesByImageOutput{
// 						FaceMatches: []types.FaceMatch{},
// 					}, nil
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uc := NewCreatePhotoUseCase(tt.repoMock, tt.rekogMock)
// 			got, err := uc.Execute(tt.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreatePhotoUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !tt.wantErr {
// 				if got == nil {
// 					t.Errorf("CreatePhotoUseCase.Execute() returned nil Photo")
// 				}
// 				if got.DeviceID != tt.input.DeviceID {
// 					t.Errorf("CreatePhotoUseCase.Execute() DeviceID = %v, want %v", got.DeviceID, tt.input.DeviceID)
// 				}
// 				if got.Recognized != tt.wantRecognized {
// 					t.Errorf("CreatePhotoUseCase.Execute() Recognized = %v, want %v", got.Recognized, tt.wantRecognized)
// 				}
// 			}
// 		})
// 	}
// }
