package integration

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/yanvic/challenge/infra/database/dynamo"
	"github.com/yanvic/challenge/internal/handler"
)

func TestMain(m *testing.M) {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("⚠️ .env not found, relying on system env")
	// }

	ctx := context.TODO()

	client, err := dynamo.InitDynamoClientTest(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dynamodb client: %v", err)
	}

	dynamo.Client = client

	os.Exit(m.Run())
}

func TestHandlerGPS_Success(t *testing.T) {
	body := `{
		"latitude": -23.55,
		"longitude": -46.63,
		"timestamp": "2025-06-21T15:00:00Z",
		"device_id": "device-1235"
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
