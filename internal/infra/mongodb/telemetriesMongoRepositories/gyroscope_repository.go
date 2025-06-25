package telemetriesMongoRepositories

import (
	"context"
	models "v3-test/internal/models/telemetriesModels"
	"v3-test/internal/repositories/telemetriesRepositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GyroscopeRepositoryMongo struct {
	collection mongo.Collection
}

func NewGyroscopeRepositoryMongo(db *mongo.Database) telemetriesRepositories.GyroscopeRepository {
	return &GyroscopeRepositoryMongo{
		collection: *db.Collection("gyroscope"),
	}
}

func (r *GyroscopeRepositoryMongo) CreateGyroscope(gyroscopeModel models.GyroscopeModel) (models.GyroscopeModel, error) {
	result, err := r.collection.InsertOne(context.Background(), gyroscopeModel)
	if err != nil {
		return models.GyroscopeModel{}, err
	}

	insertedID := result.InsertedID

	var created models.GyroscopeModel

	err = r.collection.FindOne(context.Background(), bson.M{"_id": insertedID}).Decode(&created)
	if err != nil {
		return models.GyroscopeModel{}, err
	}

	return created, nil
}
