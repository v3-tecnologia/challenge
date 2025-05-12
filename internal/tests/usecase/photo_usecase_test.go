package usecase

import (
	"errors"
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/tests/usecase/mocks"
	"v3/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePhotoUseCase_Execute(t *testing.T) {
	// Timestamp fixo para consistência
	fixedTimestamp := time.Now().Truncate(time.Second).UTC()
	fixedUnix := fixedTimestamp.Unix()

	tests := []struct {
		name           string
		input          domain.PhotoDto
		photoBytes     []byte
		setupMocks     func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService)
		wantErr        error
		validateResult func(t *testing.T, photo *domain.Photo)
	}{
		{
			name: "Successful photo creation",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", []byte("photo-data"), fixedUnix).Return("/photos/123.jpg", nil).Once()
				aws.On("ComparePhoto", "00:0a:95:9d:68:16", "/photos/123.jpg").Return(true, nil).Once()
				photo := &domain.Photo{
					ID:         "mock-id",
					DeviceID:   "00:0a:95:9d:68:16",
					FilePath:   "/photos/123.jpg",
					Recognized: true,
					Timestamp:  fixedTimestamp,
				}
				repo.On("Create", mock.MatchedBy(func(p *domain.Photo) bool {
					return p.DeviceID == "00:0a:95:9d:68:16" &&
						p.FilePath == "/photos/123.jpg" &&
						p.Recognized == true &&
						p.Timestamp == fixedTimestamp
				})).Return(photo, nil).Once()
			},
			wantErr: nil,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.NotNil(t, photo)
				assert.Equal(t, "mock-id", photo.ID)
				assert.Equal(t, "00:0a:95:9d:68:16", photo.DeviceID)
				assert.Equal(t, "/photos/123.jpg", photo.FilePath)
				assert.True(t, photo.Recognized)
				assert.Equal(t, fixedTimestamp, photo.Timestamp)
			},
		},
		{
			name: "Invalid DeviceID",
			input: domain.PhotoDto{
				DeviceID:  "",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				// No mocks needed
			},
			wantErr: domain.ErrDeviceIDPhoto,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "Zero Timestamp",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: 0,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				// No mocks needed
			},
			wantErr: domain.ErrTimestampPhoto,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "Empty Photo Bytes",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte{},
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				// No mocks needed
			},
			wantErr: domain.ErrPhotoData,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "AWS UploadPhoto error",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", []byte("photo-data"), fixedUnix).Return("", errors.New("upload failed")).Once()
			},
			wantErr: domain.ErrProcessPhotoWithAWSRekognition,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "AWS ComparePhoto error",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", []byte("photo-data"), fixedUnix).Return("/photos/123.jpg", nil).Once()
				aws.On("ComparePhoto", "00:0a:95:9d:68:16", "/photos/123.jpg").Return(false, errors.New("compare failed")).Once()
			},
			wantErr: domain.ErrProcessPhotoWithAWSRekognition,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "Repository Create error",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("photo-data"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", []byte("photo-data"), fixedUnix).Return("/photos/123.jpg", nil).Once()
				aws.On("ComparePhoto", "00:0a:95:9d:68:16", "/photos/123.jpg").Return(true, nil).Once()
				repo.On("Create", mock.AnythingOfType("*domain.Photo")).Return(nil, errors.New("database error")).Once()
			},
			wantErr: domain.ErrSavePhotoData,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "Photo size exceeds 5MB",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: make([]byte, 6*1024*1024), // 6MB
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", mock.MatchedBy(func(b []byte) bool { return len(b) == 6*1024*1024 }), fixedUnix).Return("", errors.New("photo size exceeds 5MB")).Once()
			},
			wantErr: domain.ErrProcessPhotoWithAWSRekognition,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
		{
			name: "Unsupported photo format",
			input: domain.PhotoDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Timestamp: fixedUnix,
			},
			photoBytes: []byte("invalid-format"),
			setupMocks: func(repo *mocks.MockPhotoRepository, aws *mocks.MockAWSService) {
				aws.On("UploadPhoto", "00:0a:95:9d:68:16", []byte("invalid-format"), fixedUnix).Return("", errors.New("unsupported photo format: application/octet-stream")).Once()
			},
			wantErr: domain.ErrProcessPhotoWithAWSRekognition,
			validateResult: func(t *testing.T, photo *domain.Photo) {
				assert.Nil(t, photo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializar mocks
			repo := &mocks.MockPhotoRepository{}
			awsService := &mocks.MockAWSService{}
			tt.setupMocks(repo, awsService)

			// Criar use case
			uc := usecase.NewCreatePhotoUseCase(repo, awsService)

			// Executar use case
			result, err := uc.Execute(tt.input, tt.photoBytes)

			// Validar erro
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			// Validar resultado
			tt.validateResult(t, result)

			// Verificar se os mocks foram chamados como esperado
			repo.AssertExpectations(t)
			awsService.AssertExpectations(t)
		})
	}
}
