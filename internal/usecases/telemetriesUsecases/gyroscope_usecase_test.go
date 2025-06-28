package telemetriesUsecases

import (
	"errors"
	"testing"

	dtos "v3-test/internal/dtos/telemetriesDtos"
	models "v3-test/internal/models/telemetriesModels"
)

type mockGyroscopeRepo struct {
	mockCreateFunc func(models.GyroscopeModel) (models.GyroscopeModel, error)
}

func (m *mockGyroscopeRepo) CreateGyroscope(g models.GyroscopeModel) (models.GyroscopeModel, error) {
	return m.mockCreateFunc(g)
}

func TestCreateGyroscope_Success(t *testing.T) {
	mock := &mockGyroscopeRepo{
		mockCreateFunc: func(g models.GyroscopeModel) (models.GyroscopeModel, error) {
			return g, nil
		},
	}

	usecase := NewGyroscopeUsecase(mock)

	x := 1.23
	y := 4.56
	z := 7.89

	dto := dtos.CreateGyroscopeDto{
		X: &x,
		Y: &y,
		Z: &z,
	}

	result, err := usecase.CreateGyroscope(dto)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.X == nil || result.Y == nil || result.Z == nil {
		t.Fatal("result contains nil values")
	}
	if dto.X == nil || dto.Y == nil || dto.Z == nil {
		t.Fatal("dto contains nil values")
	}

	if *result.X != *dto.X || *result.Y != *dto.Y || *result.Z != *dto.Z {
		t.Errorf("expected values (%.2f, %.2f, %.2f), got (%.2f, %.2f, %.2f)",
			*dto.X, *dto.Y, *dto.Z, *result.X, *result.Y, *result.Z)
	}
}

func TestCreateGyroscope_ErrorOnRepo(t *testing.T) {
	mock := &mockGyroscopeRepo{
		mockCreateFunc: func(g models.GyroscopeModel) (models.GyroscopeModel, error) {
			return models.GyroscopeModel{}, errors.New("mock error")
		},
	}

	usecase := NewGyroscopeUsecase(mock)

	x := 1.3
	y := 8.7
	z := 9.6

	dto := dtos.CreateGyroscopeDto{
		X: &x,
		Y: &y,
		Z: &z,
	}

	_, err := usecase.CreateGyroscope(dto)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "mock error" {
		t.Errorf("expected 'mock error', got '%s'", err.Error())
	}
}
