package commonmongorepositories

import (
	"context"
	"fmt"
	"v3-test/internal/models/commonModels"
	"v3-test/internal/repositories/commonRepositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhotoMongoRepository struct {
	collection *mongo.Collection
}

func NewPhotoMongoRepository(db *mongo.Database) commonRepositories.PhotoRepository {
	return &PhotoMongoRepository{
		collection: db.Collection("photos"),
	}
}

func (r *PhotoMongoRepository) CreatePhoto(photoModel commonModels.PhotoModel) (commonModels.PhotoModel, error) {
	fmt.Println("Creating photo in MongoDB repository")
	result, err := r.collection.InsertOne(context.Background(), photoModel)
	if err != nil {
		return commonModels.PhotoModel{}, err
	}

	insertedID := result.InsertedID

	var created commonModels.PhotoModel

	err = r.collection.FindOne(context.Background(), bson.M{"_id": insertedID}).Decode(&created)
	if err != nil {
		return commonModels.PhotoModel{}, err
	}

	return created, nil
}
