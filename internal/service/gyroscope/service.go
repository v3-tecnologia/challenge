package gyroscope

import (
	d "github.com/iamrosada0/v3/internal/domain/gyroscope"
	"github.com/iamrosada0/v3/internal/repository/gyroscope"
)

type GyroscopeService struct {
	Repo gyroscope.GyroscopeRepository
}

func NewGyroscopeService(repo gyroscope.GyroscopeRepository) *GyroscopeService {
	return &GyroscopeService{Repo: repo}
}

func (s *GyroscopeService) Create(data *d.GyroscopeData) (*d.GyroscopeData, error) {
	return s.Repo.Create(data)
}
