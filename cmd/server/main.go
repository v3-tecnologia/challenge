package main

import (
	"log"

	"go.uber.org/zap"

	"telemetry-api/internal/broker"
	"telemetry-api/internal/config"
	"telemetry-api/internal/database"
	"telemetry-api/internal/router"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to sync logger: %v", err)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("cannot load config", zap.Error(err))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("cannot connect to db", zap.Error(err))
	}

	nc, err := broker.ConnectNATS(cfg.NatsURL)
	if err != nil {
		logger.Fatal("cannot connect to nats", zap.Error(err))
	}

	r := router.SetupRouter(db, nc, logger)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}
