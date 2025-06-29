package databases

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func checkInternet() error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://www.google.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status not OK: %d", resp.StatusCode)
	}

	return nil
}

func ConnectMongo() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	fmt.Println("uri/dbName", uri, dbName)

	// TODO - Aprender a configurar variáveis de ambiente na AWS
	if uri == "" || dbName == "" {
		uri = "mongodb+srv://Pedro_Lomba:lCdQyijr8V3AhjVi@v3-test.rxg1mio.mongodb.net/"
		dbName = "v3-test"
	}

	fmt.Printf("Trying to connect to MongoDB at %s, database: %s\n", uri, dbName)

	fmt.Println("Checking internet connectivity...")
	if err := checkInternet(); err != nil {
		log.Fatalf("Internet connectivity check failed: %v", err)
	}
	fmt.Println("Internet connectivity OK")

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
