package usecase

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"v3/internal/domain"
// )

// type mockGPSRepository struct {
// 	createFunc func(d *domain.GPS) (*domain.GPS, error)
// }

// func (m *mockGPSRepository) Create(d *domain.GPS) (*domain.GPS, error) {
// 	return m.createFunc(d)
// }

// func TestCreateGPSUseCase_Execute(t *testing.T) {
// 	validInput := GPSInputDto{
// 		DeviceID:  "00:1A:2B:3C:4D:5E",
// 		Timestamp: time.Now().Unix(),
// 		Latitude:  -23.5505,
// 		Longitude: -46.6333,
// 	}

// 	tests := []struct {
// 		name     string
// 		input    GPSInputDto
// 		repoMock *mockGPSRepository
// 		wantErr  bool
// 	}{
// 		{
// 			name:  "valid input",
// 			input: validInput,
// 			repoMock: &mockGPSRepository{
// 				createFunc: func(d *domain.GPS) (*domain.GPS, error) {
// 					return d, nil
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "invalid input (invalid MAC)",
// 			input: GPSInputDto{
// 				DeviceID:  "invalid",
// 				Timestamp: time.Now().Unix(),
// 				Latitude:  -23.5505,
// 				Longitude: -46.6333,
// 			},
// 			repoMock: &mockGPSRepository{
// 				createFunc: func(d *domain.GPS) (*domain.GPS, error) {
// 					return nil, errors.New("should not be called")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name:  "repository error",
// 			input: validInput,
// 			repoMock: &mockGPSRepository{
// 				createFunc: func(d *domain.GPS) (*domain.GPS, error) {
// 					return nil, errors.New("database error")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uc := NewCreateGPSUseCase(tt.repoMock)
// 			got, err := uc.Execute(tt.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateGPSUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !tt.wantErr {
// 				if got == nil {
// 					t.Errorf("CreateGPSUseCase.Execute() returned nil GPS")
// 				}
// 				if got.DeviceID != tt.input.DeviceID {
// 					t.Errorf("CreateGPSUseCase.Execute() DeviceID = %v, want %v", got.DeviceID, tt.input.DeviceID)
// 				}
// 			}
// 		})
// 	}
// }
