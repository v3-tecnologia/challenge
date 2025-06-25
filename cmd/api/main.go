package main

import (
	"context"
	"log"
	"net/http"

	"github.com/yanvic/challenge/infra/database/dynamo"
	"github.com/yanvic/challenge/internal/handler"
	"github.com/yanvic/challenge/internal/queue"
)

func main() {
	ctx := context.TODO()

	client, err := dynamo.InitDynamoClient(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dynamodb client: %v", err)
	}
	dynamo.Client = client

	// ✅ Espera o DynamoDB estar pronto antes de criar as tabelas
	if err := dynamo.WaitForDynamoReady(ctx, client); err != nil {
		log.Fatalf("erro esperando DynamoDB: %v", err)
	}

	tables := []string{"GyroscopeTable", "GPSTable", "PhotoTable", "PhotoAnalysisTable"}

	for _, t := range tables {
		if err := dynamo.EnsureTable(ctx, client, t); err != nil {
			log.Fatalf("erro ao garantir tabela %s: %v", t, err)
		}
	}

	queue.InitNATS()

	go queue.StartImageConsumer()

	http.HandleFunc("/telemetry/gyroscope", handler.HandlerGyroscope)
	http.HandleFunc("/telemetry/gps", handler.HandlerGPS)
	http.HandleFunc("/telemetry/photo", handler.HandlerPhoto)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
