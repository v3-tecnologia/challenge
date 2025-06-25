package bootstrap

import (
	"context"
	"v3-test/internal/config/databases"

	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

func SetupDatabase() *mongo.Database {
	db := databases.ConnectMongo()
	return db
}
