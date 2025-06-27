package main

import (
	_ "challenge-v3/docs" // Import para o Swagger
	"challenge-v3/handlers"
	"challenge-v3/messaging"
	"challenge-v3/metrics"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           API de Telemetria de Frota
// @version         1.0
// @description     Esta é a API para ingestão de dados de telemetria do Desafio Cloud.
// @host      localhost:8080
// @BasePath  /
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	godotenv.Load()
	slog.Info("Iniciando a API...")

	natsURL := os.Getenv("NATS_URL")
	nc, err := messaging.ConnectNATS(natsURL)
	if err != nil {
		slog.Error("Falha ao conectar ao NATS", "error", err)
		os.Exit(1)
	}
	defer nc.Close()

	js, err := messaging.SetupJetStream(nc)
	if err != nil {
		slog.Error("Falha ao configurar o JetStream", "error", err)
		os.Exit(1)
	}

	// A API só precisa da dependência do NATS para publicar mensagens.
	// As outras dependências (DB, Rekognition) são usadas apenas pelo Worker.
	api := handlers.NewAPI(nil, nil, js)

	router := http.NewServeMux()

	router.Handle("/telemetry/gyroscope",
		handlers.RateLimiterMiddleware(handlers.AuthenticationMiddleware(metrics.PrometheusMiddleware(http.HandlerFunc(api.HandleGyroscope)))))

	router.Handle("/telemetry/gps",
		handlers.RateLimiterMiddleware(handlers.AuthenticationMiddleware(metrics.PrometheusMiddleware(http.HandlerFunc(api.HandleGPS)))))

	router.Handle("/telemetry/photo",
		handlers.RateLimiterMiddleware(handlers.AuthenticationMiddleware(metrics.PrometheusMiddleware(http.HandlerFunc(api.HandlePhoto)))))

	router.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	metricsRouter := http.NewServeMux()
	metricsRouter.Handle("/metrics", metrics.MetricsHandler())
	go func() {
		slog.Info("Servidor de Métricas da API iniciado", "porta", ":8081")
		if err := http.ListenAndServe(":8081", metricsRouter); err != nil {
			slog.Error("Servidor de Métricas da API falhou", "error", err)
		}
	}()

	slog.Info("Servidor da API iniciado", "porta", ":8080", "swagger", "http://localhost:8080/swagger/index.html")
	if err := http.ListenAndServe(":8080", router); err != nil {
		slog.Error("servidor http da API falhou", "error", err)
		os.Exit(1)
	}
}
