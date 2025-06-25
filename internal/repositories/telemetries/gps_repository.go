package repositories

import models "v3-test/internal/models/telemetries"

type GpsRepository interface {
	InsertGps(gps models.GpsModel) error
}
