package handlers

import (
	"bytes"
	"challenge-v3/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockNATSJetStream struct {
	nats.JetStreamContext
	mock.Mock
}

func (m *MockNATSJetStream) Publish(subj string, data []byte, opts ...nats.PubOpt) (*nats.PubAck, error) {
	args := m.Called(subj, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*nats.PubAck), args.Error(1)
}

func float64Ptr(f float64) *float64 { return &f }

func TestAuthenticationMiddleware(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Log("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema/CI.")
	}

	apiKey := os.Getenv("API_KEY")
	require.NotEmpty(t, apiKey, "API_KEY não pode ser vazia no .env")

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Acesso Permitido"))
	})

	protectedHandler := AuthenticationMiddleware(dummyHandler)

	t.Run("falha - sem chave de api", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		protectedHandler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Deveria retornar 401 sem a chave")
	})

	t.Run("falha - com chave de api errada", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("X-API-Key", "chave-errada-123")
		rr := httptest.NewRecorder()
		protectedHandler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code, "Deveria retornar 401 com a chave errada")
	})

	t.Run("sucesso - com chave de api correta", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("X-API-Key", apiKey)
		rr := httptest.NewRecorder()
		protectedHandler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, "Deveria retornar 200 com a chave correta")
	})
}

func TestHandleGyroscope_Async(t *testing.T) {
	mockJS := new(MockNATSJetStream)
	api := NewAPI(nil, nil, mockJS)
	apiKey := os.Getenv("API_KEY")

	testData := models.GyroscopeData{
		DeviceID:  "gyro-test-async",
		X:         float64Ptr(1),
		Y:         float64Ptr(2),
		Z:         float64Ptr(3),
		Timestamp: time.Now(),
	}
	payloadBytes, err := json.Marshal(testData)
	require.NoError(t, err)
	mockJS.On("Publish", "telemetry.gyroscope", payloadBytes).Return(&nats.PubAck{}, nil)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gyroscope", bytes.NewBuffer(payloadBytes))
	req.Header.Set("X-API-Key", apiKey)
	rr := httptest.NewRecorder()
	api.HandleGyroscope(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	mockJS.AssertExpectations(t)
}

func TestHandleGPS_Async(t *testing.T) {
	mockJS := new(MockNATSJetStream)
	api := NewAPI(nil, nil, mockJS)
	apiKey := os.Getenv("API_KEY")

	testData := models.GPSData{
		DeviceID:  "gps-test-async",
		Latitude:  float64Ptr(10),
		Longitude: float64Ptr(20),
		Timestamp: time.Now(),
	}
	payloadBytes, err := json.Marshal(testData)
	require.NoError(t, err)
	mockJS.On("Publish", "telemetry.gps", payloadBytes).Return(&nats.PubAck{}, nil)

	req := httptest.NewRequest(http.MethodPost, "/telemetry/gps", bytes.NewBuffer(payloadBytes))
	req.Header.Set("X-API-Key", apiKey)
	rr := httptest.NewRecorder()
	api.HandleGPS(rr, req)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	mockJS.AssertExpectations(t)
}

func TestHandlePhoto_Async(t *testing.T) {
	apiKey := os.Getenv("API_KEY")

	t.Run("sucesso - publica mensagem na fila", func(t *testing.T) {
		mockJS := new(MockNATSJetStream)
		api := NewAPI(nil, nil, mockJS)

		requestData := models.PhotoRequest{
			DeviceID:  "photo-test-async",
			Photo:     "aW1hZ2VtLWRhdGE=",
			Timestamp: time.Now(),
		}
		requestPayloadBytes, err := json.Marshal(requestData)
		require.NoError(t, err)

		expectedMessageData := models.PhotoData{
			DeviceID:   requestData.DeviceID,
			Photo:      requestData.Photo,
			Timestamp:  requestData.Timestamp,
			Recognized: false,
		}
		expectedMessageBytes, err := json.Marshal(expectedMessageData)
		require.NoError(t, err)

		mockJS.On("Publish", "telemetry.photo", expectedMessageBytes).Return(&nats.PubAck{}, nil)

		req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", bytes.NewBuffer(requestPayloadBytes))
		req.Header.Set("X-API-Key", apiKey)
		rr := httptest.NewRecorder()
		api.HandlePhoto(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
		mockJS.AssertExpectations(t)
	})

	t.Run("falha - dados de validação inválidos", func(t *testing.T) {
		mockJS := new(MockNATSJetStream)
		api := NewAPI(nil, nil, mockJS)

		requestData := models.PhotoRequest{DeviceID: "photo-test-invalid"} // Faltam campos
		payloadBytes, err := json.Marshal(requestData)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/telemetry/photo", bytes.NewBuffer(payloadBytes))
		req.Header.Set("X-API-Key", apiKey)
		rr := httptest.NewRecorder()
		api.HandlePhoto(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockJS.AssertNotCalled(t, "Publish", mock.Anything)
	})
}
