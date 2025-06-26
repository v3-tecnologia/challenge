package bootstrap

import (
	commonmongorepositories "v3-test/internal/infra/mongodb/commonMongoRepositories"
	"v3-test/internal/infra/mongodb/telemetriesMongoRepositories"
	"v3-test/internal/repositories/commonRepositories"
	"v3-test/internal/repositories/telemetriesRepositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	GpsRepo       telemetriesRepositories.GpsRepository
	GyroscopeRepo telemetriesRepositories.GyroscopeRepository
	PhotoRepo     commonRepositories.PhotoRepository
}

func SetupRepositories(db *mongo.Database) Repositories {
	return Repositories{
		GpsRepo:       telemetriesMongoRepositories.NewGpsRepositoryMongo(db),
		GyroscopeRepo: telemetriesMongoRepositories.NewGyroscopeRepositoryMongo(db),
		PhotoRepo:     commonmongorepositories.NewPhotoMongoRepository(db),
	}
}
