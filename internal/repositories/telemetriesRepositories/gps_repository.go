package telemetriesRepositories

import models "v3-test/internal/models/telemetriesModels"

type GpsRepository interface {
	CreateGps(gps models.GpsModel) (models.GpsModel, error)
}
