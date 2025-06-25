package mongoRepositories

import (
	"context"
	models "v3-test/internal/models/telemetries"
	"v3-test/internal/repositories/telemetries"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *gpsRepositoryMongo) CreateGps(gps models.GpsModel) (models.GpsModel, error) {
	result, err := r.collection.InsertOne(context.Background(), gps)
	if err != nil {
		return models.GpsModel{}, err
	}

	insertedID := result.InsertedID

	var created models.GpsModel

	err = r.collection.FindOne(context.Background(), bson.M{"_id": insertedID}).Decode(&created)

	if err != nil {
		return models.GpsModel{}, err
	}

	return created, nil
}
