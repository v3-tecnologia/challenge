package usecase

import (
	"errors"
	"math"
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/tests/usecase/mocks"
	"v3/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateGPSUseCase_Execute(t *testing.T) {
	// Timestamp fixo para evitar variações
	fixedTimestamp := time.Now().Truncate(time.Second).UTC()
	fixedUnix := fixedTimestamp.Unix()

	tests := []struct {
		name           string
		input          domain.GPSDto
		setupMocks     func(repo *mocks.MockGPSRepository)
		wantErr        error
		validateResult func(t *testing.T, gps *domain.GPS)
	}{
		{
			name: "Successful GPS creation",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMocks: func(repo *mocks.MockGPSRepository) {
				gps := &domain.GPS{
					ID:        "mock-id",
					DeviceID:  "00:0a:95:9d:68:16",
					Latitude:  -23.5505,
					Longitude: -46.6333,
					Timestamp: fixedTimestamp,
				}
				repo.On("Create", mock.MatchedBy(func(g *domain.GPS) bool {
					return g.DeviceID == "00:0a:95:9d:68:16" &&
						g.Latitude == -23.5505 &&
						g.Longitude == -46.6333 &&
						g.Timestamp.Equal(fixedTimestamp)
				})).Return(gps, nil).Once()
			},
			wantErr: nil,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.NotNil(t, gps)
				assert.Equal(t, "mock-id", gps.ID)
				assert.Equal(t, "00:0a:95:9d:68:16", gps.DeviceID)
				assert.Equal(t, -23.5505, gps.Latitude)
				assert.Equal(t, -46.6333, gps.Longitude)
				assert.Equal(t, fixedTimestamp, gps.Timestamp)
			},
		},
		{
			name: "Invalid DeviceID",
			input: domain.GPSDto{
				DeviceID:  "",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMocks: func(repo *mocks.MockGPSRepository) {
				// No mocks needed
			},
			wantErr: domain.ErrDeviceIDGPS,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Zero Timestamp",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: 0,
			},
			setupMocks: func(repo *mocks.MockGPSRepository) {
				// No mocks needed
			},
			wantErr: domain.ErrTimestampGPS,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Invalid Latitude (NaN)",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  math.NaN(),
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMocks: func(repo *mocks.MockGPSRepository) {
				// No mocks needed
			},
			wantErr: domain.ErrInvalidGPSValues,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
		{
			name: "Repository error",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMocks: func(repo *mocks.MockGPSRepository) {
				repo.On("Create", mock.AnythingOfType("*domain.GPS")).Return(nil, errors.New("database error")).Once()
			},
			wantErr: domain.ErrSaveGPSData,
			validateResult: func(t *testing.T, gps *domain.GPS) {
				assert.Nil(t, gps)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializar mock
			repo := &mocks.MockGPSRepository{}
			tt.setupMocks(repo)

			// Criar use case
			uc := usecase.NewCreateGPSUseCase(repo)

			// Executar use case
			result, err := uc.Execute(tt.input)

			// Validar erro
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			// Validar resultado
			tt.validateResult(t, result)

			// Verificar se o mock foi chamado como esperado
			repo.AssertExpectations(t)
		})
	}
}
