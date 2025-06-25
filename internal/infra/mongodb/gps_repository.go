package mongoRepositories

import (
	"context"
	models "v3-test/internal/models/telemetries"
	"v3-test/internal/repositories/telemetries"

	"go.mongodb.org/mongo-driver/mongo"
)

type gpsRepositoryMongo struct {
	collection *mongo.Collection
}

func NewGpsRepositoryMongo(db *mongo.Database) telemetries.GpsRepository {
	return &gpsRepositoryMongo{
		collection: db.Collection("gps"),
	}
}

func (r *gpsRepositoryMongo) CreateGps(gps models.GpsModel) error {
	_, err := r.collection.InsertOne(context.Background(), gps)
	return err
}
