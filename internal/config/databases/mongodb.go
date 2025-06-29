package databases

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	fmt.Println("uri/dbName", uri, dbName)

	// TODO - Aprender a configurar variáveis de ambiente na AWS
	if uri == "" || dbName == "" {
		uri = "mongodb+srv://Pedro_Lomba:lCdQyijr8V3AhjVi@v3-test.rxg1mio.mongodb.net/"
		dbName = "v3-test"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("MongoDB connected successfully!")

	MongoClient = client
	return client.Database(dbName)
}
