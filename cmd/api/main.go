// @title Challenge API
// @version 1.0
// @description API for telemetry challenge
// @host localhost:8080
// @BasePath /
package main

import (
	"context"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yanvic/challenge/docs"
	"github.com/yanvic/challenge/infra/database/dynamo"
	"github.com/yanvic/challenge/internal/handler"
	"github.com/yanvic/challenge/internal/queue"
)

// @Summary Envia dados do GPS
// @Description Envia latitude, longitude e timestamp
// @Accept  json
// @Produce  json
// @Param gps body entity.GPS true "Dados do GPS"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /telemetry/gps [post]

func main() {
	ctx := context.TODO()

	client, err := dynamo.InitDynamoClient(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dynamodb client: %v", err)
	}
	dynamo.Client = client

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
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
