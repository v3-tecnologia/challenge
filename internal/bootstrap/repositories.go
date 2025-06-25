package bootstrap

import (
	"v3-test/internal/infra/mongodb/telemetriesMongoRepositories"
	"v3-test/internal/repositories/telemetriesRepositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	GpsRepo       telemetriesRepositories.GpsRepository
	GyroscopeRepo telemetriesRepositories.GyroscopeRepository
}

func SetupRepositories(db *mongo.Database) Repositories {
	return Repositories{
		GpsRepo:       telemetriesMongoRepositories.NewGpsRepositoryMongo(db),
		GyroscopeRepo: telemetriesMongoRepositories.NewGyroscopeRepositoryMongo(db),
	}
}
