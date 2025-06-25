package bootstrap

import (
	mongoRepositories "v3-test/internal/infra/mongodb"
	"v3-test/internal/repositories/telemetries"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	GpsRepo telemetries.GpsRepository
}

func SetupRepositories(db *mongo.Database) Repositories {
	return Repositories{
		GpsRepo: mongoRepositories.NewGpsRepositoryMongo(db),
	}
}
