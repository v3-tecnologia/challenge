package integration

import (
	"context"
	"log"
	"os"
	"testing"
	"v3-test/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app *gin.Engine
var db *mongo.Database

// TestMain roda antes de todos os testes no pacote
func TestMain(m *testing.M) {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	db = client.Database("test-db")
	app = bootstrap.BuildApp(db)

	code := m.Run()

	db.Drop(context.TODO())

	os.Exit(code)
}

// Limpa as coleções antes de cada teste
func ClearDatabase(t *testing.T) {
	collections := []string{"gps", "gyroscope", "photos"}

	for _, name := range collections {
		err := db.Collection(name).Drop(context.Background())
		if err != nil {
			t.Fatalf("Erro ao limpar a collection %s: %v", name, err)
		}
	}
}
