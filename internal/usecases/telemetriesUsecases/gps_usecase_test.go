package telemetriesUsecases

import (
	"errors"
	"testing"

	"v3-test/internal/dtos/telemetriesDtos"
	"v3-test/internal/models/telemetriesModels"
)

type mockGpsRepo struct {
	mockCreateFunc func(telemetriesModels.GpsModel) (telemetriesModels.GpsModel, error)
}

func (m *mockGpsRepo) CreateGps(g telemetriesModels.GpsModel) (telemetriesModels.GpsModel, error) {
	return m.mockCreateFunc(g)
}

func TestCreateGps_Success(t *testing.T) {
	mock := &mockGpsRepo{
		mockCreateFunc: func(g telemetriesModels.GpsModel) (telemetriesModels.GpsModel, error) {
			return g, nil
		},
	}

	usecase := NewGpsUsecase(mock)

	latitude := 1.1
	longitude := 2.2

	dto := telemetriesDtos.CreateGpsDto{
		Latitude:  &latitude,
		Longitude: &longitude,
	}

	result, err := usecase.CreateGps(dto)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Latitude != &latitude || result.Longitude != &longitude {
		t.Errorf("expected coordinates to match input, got %f, %f", *result.Latitude, *result.Longitude)
	}
}

func TestCreateGps_ErrorOnRepo(t *testing.T) {
	mock := &mockGpsRepo{
		mockCreateFunc: func(g telemetriesModels.GpsModel) (telemetriesModels.GpsModel, error) {
			return telemetriesModels.GpsModel{}, errors.New("mock error")
		},
	}

	usecase := NewGpsUsecase(mock)

	latitude := 1.1
	longitude := 2.2

	dto := telemetriesDtos.CreateGpsDto{
		Latitude:  &latitude,
		Longitude: &longitude,
	}

	_, err := usecase.CreateGps(dto)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "mock error" {
		t.Errorf("expected 'mock error', got '%s'", err.Error())
	}
}
