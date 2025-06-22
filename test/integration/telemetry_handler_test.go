package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yanvic/challenge/internal/handler"
)

func TestHandlerGPS_Success(t *testing.T) {
	body := `{
		"latitude": -23.55,
		"longitude": -46.63,
		"timestamp": "2025-06-21T15:00:00Z",
		"device_id": "device-123"
	}`

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.HandlerGPS(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

func TestHandlerGPS_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", strings.NewReader("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.HandlerGPS(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}
