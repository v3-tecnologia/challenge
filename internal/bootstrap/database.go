package bootstrap

import (
	"v3-test/internal/config/databases"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupDatabase() *mongo.Database {
	db := databases.ConnectMongo()
	return db
}
