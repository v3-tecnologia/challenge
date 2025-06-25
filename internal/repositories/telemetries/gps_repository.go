package telemetries

import models "v3-test/internal/models/telemetries"

type GpsRepository interface {
	CreateGps(gps models.GpsModel) (models.GpsModel, error)
}
