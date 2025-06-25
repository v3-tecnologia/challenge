package mongodb

import (
	"context"
	"time"
	models "v3-test/internal/models/telemetries"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoGpsRepository struct {
	collection *mongo.Collection
}

func NewMongoGpsRepository(col *mongo.Collection) *MongoGpsRepository {
	return &MongoGpsRepository{collection: col}
}

func (r *MongoGpsRepository) InsertGps(gps models.GpsModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, gps)
	return err
}
