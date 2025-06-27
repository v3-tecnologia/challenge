package handlers

import (
	"challenge-v3/models"
	"challenge-v3/services"
	"challenge-v3/storage"
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
	"golang.org/x/time/rate"
)

type API struct {
	db            storage.Storage
	photoAnalyzer services.PhotoAnalyzer
	natsJS        nats.JetStreamContext
}

func NewAPI(db storage.Storage, pa services.PhotoAnalyzer, js nats.JetStreamContext) *API {
	return &API{
		db:            db,
		photoAnalyzer: pa,
		natsJS:        js,
	}
}

func SendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{Message: message})
}

var clients = make(map[string]*rate.Limiter)
var mu sync.Mutex

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			SendJSONError(w, "Erro ao obter endereço de IP", http.StatusInternalServerError)
			return
		}

		mu.Lock()

		if _, found := clients[ip]; !found {
			clients[ip] = rate.NewLimiter(rate.Limit(5), 10)
		}

		limiter := clients[ip]
		mu.Unlock()

		if !limiter.Allow() {
			SendJSONError(w, "Você atingiu o limite de requisições", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedApiKey := os.Getenv("API_KEY")
		if expectedApiKey == "" {
			slog.Error("API_KEY não está configurada no ambiente do servidor")
			SendJSONError(w, "Erro de configuração interna", http.StatusInternalServerError)
			return
		}

		receivedApiKey := r.Header.Get("X-API-Key")

		if receivedApiKey != expectedApiKey {
			slog.Warn("tentativa de acesso não autorizado", "remote_addr", r.RemoteAddr)
			SendJSONError(w, "Acesso não autorizado", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *API) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/telemetry/gyroscope", a.HandleGyroscope)
	mux.HandleFunc("/telemetry/gps", a.HandleGPS)
	mux.HandleFunc("/telemetry/photo", a.HandlePhoto)
}

// HandleGyroscope recebe e enfileira uma telemetria de giroscópio
// @Summary      Enfileira dados de telemetria de giroscópio
// @Description  Recebe um payload JSON com os dados do giroscópio, valida, e publica em uma fila NATS para processamento assíncrono.
// @Tags         Telemetry
// @Accept       json
// @Produce      json
// @Param        gyroscope   body      models.GyroscopeData  true  "Dados do Giroscópio"
// @Success      202  {object}  map[string]string
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /telemetry/gyroscope [post]
func (a *API) HandleGyroscope(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendJSONError(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var data models.GyroscopeData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		SendJSONError(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}
	if err := data.Validate(); err != nil {
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msgData, err := json.Marshal(data)
	if err != nil {
		SendJSONError(w, "Erro interno ao preparar a mensagem", http.StatusInternalServerError)
		return
	}

	_, err = a.natsJS.Publish("telemetry.gyroscope", msgData)
	if err != nil {
		slog.Error("Falha ao publicar mensagem no NATS", "topic", "telemetry.gyroscope", "error", err)
		SendJSONError(w, "Erro interno ao enviar dados para processamento", http.StatusInternalServerError)
		return
	}

	slog.Info("Mensagem publicada com sucesso", "topic", "telemetry.gps", "device_id", data.DeviceID)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Dados de giroscópio recebidos e enfileirados."})
}

// HandleGPS recebe e enfileira uma telemetria de GPS
// @Summary      Enfileira dados de telemetria de GPS
// @Description  Recebe um payload JSON com os dados de GPS, valida, e publica em uma fila NATS para processamento assíncrono.
// @Tags         Telemetry
// @Accept       json
// @Produce      json
// @Param        gps   body      models.GPSData  true  "Dados de GPS"
// @Success      202  {object}  map[string]string
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /telemetry/gps [post]
func (a *API) HandleGPS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendJSONError(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var data models.GPSData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		SendJSONError(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}
	if err := data.Validate(); err != nil {
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msgData, err := json.Marshal(data)
	if err != nil {
		SendJSONError(w, "Erro interno ao preparar a mensagem", http.StatusInternalServerError)
		return
	}

	_, err = a.natsJS.Publish("telemetry.gps", msgData)
	if err != nil {
		slog.Error("Falha ao publicar mensagem no NATS", "topic", "telemetry.gps", "error", err)
		SendJSONError(w, "Erro interno ao enviar dados para processamento", http.StatusInternalServerError)
		return
	}

	slog.Info("Mensagem publicada com sucesso", "topic", "telemetry.gps", "device_id", data.DeviceID)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Dados de GPS recebidos e enfileirados."})
}

// HandlePhoto recebe e enfileira uma telemetria de foto
// @Summary      Enfileira dados de telemetria de foto
// @Description  Recebe um payload JSON com os dados da foto, valida, e publica em uma fila NATS para processamento assíncrono.
// @Tags         Telemetry
// @Accept       json
// @Produce      json
// @Param        photo   body      models.PhotoRequest  true  "Dados da Foto a serem enviados"
// @Success      202  {object}  map[string]string
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /telemetry/photo [post]
func (a *API) HandlePhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		SendJSONError(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var requestData models.PhotoRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendJSONError(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	dataToPublish := models.PhotoData{
		DeviceID:  requestData.DeviceID,
		Photo:     requestData.Photo,
		Timestamp: requestData.Timestamp,
	}

	if err := dataToPublish.Validate(); err != nil {
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msgData, err := json.Marshal(dataToPublish)
	if err != nil {
		SendJSONError(w, "Erro interno ao preparar a mensagem", http.StatusInternalServerError)
		return
	}

	_, err = a.natsJS.Publish("telemetry.photo", msgData)
	if err != nil {
		slog.Error("Falha ao publicar mensagem no NATS", "topic", "telemetry.photo", "error", err)
		SendJSONError(w, "Erro interno ao enviar dados para processamento", http.StatusInternalServerError)
		return
	}

	slog.Info("Mensagem publicada com sucesso", "topic", "telemetry.photo", "device_id", dataToPublish.DeviceID)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Dados da foto recebidos e enfileirados para processamento."})
}
