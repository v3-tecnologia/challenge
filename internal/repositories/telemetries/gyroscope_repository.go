package telemetries

import models "v3-test/internal/models/telemetries"

type GyroscopeRepository interface {
	CreateGyroscope(gyroscopeModel models.GyroscopeModel) (models.GyroscopeModel, error)
}
