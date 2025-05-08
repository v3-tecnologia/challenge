package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"Request: %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		log.Printf(
			"Response: %s %s - completed in %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
