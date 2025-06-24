package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	handler "github.com/dryingcore/v3-challenge/internal/adapter/handler/http"
	"github.com/dryingcore/v3-challenge/internal/config"
	corequeue "github.com/dryingcore/v3-challenge/internal/core/ports/queue"
	"github.com/dryingcore/v3-challenge/internal/core/usecase"
	"github.com/dryingcore/v3-challenge/pkg/queue"
)

func main() {
	config.Load()

	var publisher corequeue.Publisher
	pub, err := queue.NewRealPublisher(config.NATSUrl)
	if err != nil {
		log.Printf("⚠️  Unable to connect on NATS (%s), using DummyPublisher: %v", config.NATSUrl, err)
		publisher = queue.NewFakePublisher()
	} else {
		log.Printf("✅ Connected on NATS using %s", config.NATSUrl)
		publisher = pub
	}

	_handler := handler.NewTelemetryHandler(
		usecase.NewGyroscopeUC(publisher),
		usecase.NewGPSUseCase(publisher),
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/telemetry/gyroscope", _handler.HandleGyroscope)
	r.Post("/telemetry/gps", _handler.HandleGPS)

	log.Println("🚀 Server running at port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
