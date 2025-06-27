package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var HTTPRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP recebidas.",
	},
	[]string{"method", "path", "status_code"},
)

var HTTPRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duração das requisições HTTP em segundos.",
	},
	[]string{"method", "path"},
)

var NatsMessagesProcessed = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "nats_messages_processed_total",
		Help: "Total de mensagens NATS processadas.",
	},
	[]string{"subject", "status"},
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(rw.statusCode)

		HTTPRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		HTTPRequestsTotal.WithLabelValues(r.Method, r.URL.Path, statusCode).Inc()
	})
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
