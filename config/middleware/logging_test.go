package middleware

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	var logOutput strings.Builder
	log.SetOutput(&logOutput)
	defer log.SetOutput(io.Discard)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot) // 418 :)
		w.Write([]byte("I'm a teapot"))
	})

	testHandler := LoggingMiddleware(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	testHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTeapot {
		t.Errorf("esperado status %d, recebido %d", http.StatusTeapot, rr.Code)
	}

	logs := logOutput.String()
	if !strings.Contains(logs, "Started GET /test") {
		t.Errorf("log não contém início da requisição: %s", logs)
	}
	if !strings.Contains(logs, "Completed /test") {
		t.Errorf("log não contém fim da requisição: %s", logs)
	}
}
