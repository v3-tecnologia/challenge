package telemetriesRepositories

import models "v3-test/internal/models/telemetriesModels"

type GyroscopeRepository interface {
	CreateGyroscope(gyroscopeModel models.GyroscopeModel) (models.GyroscopeModel, error)
}
