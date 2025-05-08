package gyroscope

import (
	"time"

	"github.com/iamrosada0/v3/internal/domain/gyroscope"
)

func (uc *GyroscopeUseCase) Create(deviceID string, x, y, z float64, timestamp time.Time) (*gyroscope.GyroscopeData, error) {
	data, err := gyroscope.NewGyroscopeData(deviceID, x, y, z, timestamp)
	if err != nil {
		return nil, err
	}
	return uc.Repo.Create(data)
}
