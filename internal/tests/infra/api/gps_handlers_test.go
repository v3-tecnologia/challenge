package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"v3/internal/domain"
	"v3/internal/infra/api"
	"v3/internal/tests/infra/api/mocks" // Atualizado para infra/api
	"v3/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(useCase usecase.GPSUseCase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handlers := api.NewGPSHandlers(useCase)
	handlers.SetupRoutes(router)
	return router
}

func TestGPSHandlers_CreateGPSHandler(t *testing.T) {
	// Timestamp fixo para consistência
	fixedTimestamp := time.Now().Truncate(time.Second).UTC()
	fixedUnix := fixedTimestamp.Unix()

	tests := []struct {
		name           string
		input          interface{}
		setupMock      func(useCase *mocks.MockCreateGPSUseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Successful GPS creation",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMock: func(useCase *mocks.MockCreateGPSUseCase) {
				gps := &domain.GPS{
					ID:        "mock-id",
					DeviceID:  "00:0a:95:9d:68:16",
					Latitude:  -23.5505,
					Longitude: -46.6333,
					Timestamp: fixedTimestamp,
				}
				useCase.On("Execute", mock.MatchedBy(func(dto domain.GPSDto) bool {
					return dto.DeviceID == "00:0a:95:9d:68:16" &&
						dto.Latitude == -23.5505 &&
						dto.Longitude == -46.6333 &&
						dto.Timestamp == fixedUnix
				})).Return(gps, nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":        "mock-id",
				"device_id": "00:0a:95:9d:68:16",
				"latitude":  -23.5505,
				"longitude": -46.6333,
				"timestamp": fixedTimestamp.Format(time.RFC3339),
			},
		},

		{
			name: "Repository error",
			input: domain.GPSDto{
				DeviceID:  "00:0a:95:9d:68:16",
				Latitude:  -23.5505,
				Longitude: -46.6333,
				Timestamp: fixedUnix,
			},
			setupMock: func(useCase *mocks.MockCreateGPSUseCase) {
				useCase.On("Execute", mock.Anything).Return(nil, domain.ErrSaveGPSData).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": domain.ErrSaveGPSData.Error(), // Ajustar se souber a string exata
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializar mock
			useCase := &mocks.MockCreateGPSUseCase{}
			tt.setupMock(useCase)

			// Configurar router
			router := setupRouter(useCase)

			// Criar requisição
			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/api/telemetry/gps", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Executar handler
			router.ServeHTTP(w, req)

			// Validar status
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Validar corpo da resposta
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err, "Failed to unmarshal response body")
			for k, v := range tt.expectedBody {
				assert.Equal(t, v, response[k], "Field %s mismatch", k)
			}

			// Verificar se o mock foi chamado como esperado
			useCase.AssertExpectations(t)
		})
	}
}
