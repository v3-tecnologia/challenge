package usecase

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"v3/internal/domain"
// )

// type mockGyroscopeRepository struct {
// 	createFunc func(d *domain.Gyroscope) (*domain.Gyroscope, error)
// }

// func (m *mockGyroscopeRepository) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {
// 	return m.createFunc(d)
// }

// func TestCreateGyroscopeUseCase_Execute(t *testing.T) {
// 	validInput := GyroscopeInputDto{
// 		DeviceID:  "00:1A:2B:3C:4D:5E",
// 		Timestamp: time.Now().Unix(),
// 		X:         1.0,
// 		Y:         2.0,
// 		Z:         3.0,
// 	}

// 	tests := []struct {
// 		name     string
// 		input    GyroscopeInputDto
// 		repoMock *mockGyroscopeRepository
// 		wantErr  bool
// 	}{
// 		{
// 			name:  "valid input",
// 			input: validInput,
// 			repoMock: &mockGyroscopeRepository{
// 				createFunc: func(d *domain.Gyroscope) (*domain.Gyroscope, error) {
// 					return d, nil
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "invalid input (invalid MAC)",
// 			input: GyroscopeInputDto{
// 				DeviceID:  "invalid",
// 				Timestamp: time.Now().Unix(),
// 				X:         1.0,
// 				Y:         2.0,
// 				Z:         3.0,
// 			},
// 			repoMock: &mockGyroscopeRepository{
// 				createFunc: func(d *domain.Gyroscope) (*domain.Gyroscope, error) {
// 					return nil, errors.New("should not be called")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name:  "repository error",
// 			input: validInput,
// 			repoMock: &mockGyroscopeRepository{
// 				createFunc: func(d *domain.Gyroscope) (*domain.Gyroscope, error) {
// 					return nil, errors.New("database error")
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			uc := NewCreateGyroscopeUseCase(tt.repoMock)
// 			got, err := uc.Execute(tt.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateGyroscopeUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !tt.wantErr {
// 				if got == nil {
// 					t.Errorf("CreateGyroscopeUseCase.Execute() returned nil Gyroscope")
// 				}
// 				if got.DeviceID != tt.input.DeviceID {
// 					t.Errorf("CreateGyroscopeUseCase.Execute() DeviceID = %v, want %v", got.DeviceID, tt.input.DeviceID)
// 				}
// 			}
// 		})
// 	}
// }
