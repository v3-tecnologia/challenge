package service

import (
	d "github.com/iamrosada0/v3/internal/domain/gyroscope"

	"github.com/stretchr/testify/mock"
)

type MockGyroscopeRepository struct {
	mock.Mock
}

func (m *MockGyroscopeRepository) Create(data *d.GyroscopeData) (*d.GyroscopeData, error) {
	args := m.Called(data)
	return args.Get(0).(*d.GyroscopeData), args.Error(1)
}
