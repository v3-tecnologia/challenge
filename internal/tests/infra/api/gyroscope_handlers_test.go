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
	"v3/internal/tests/infra/api/mocks"
	"v3/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupGyroscopeRouter(useCase usecase.GyroscopeUseCase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handlers := api.NewGyroscopeHandlers(useCase)
	handlers.SetupRoutes(router)
	return router
}

func TestGyroscopeHandlers_CreateGyroscopeHandler(t *testing.T) {
	// Timestamp fixo para consistência
	fixedTimestamp := time.Now().Truncate(time.Second).UTC()
	fixedUnix := fixedTimestamp.Unix()

	tests := []struct {
		name           string
		input          interface{}
		setupMock      func(useCase *mocks.MockCreateGyroscopeUseCase)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Successful Gyroscope creation",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				X:         1.0,
				Y:         2.0,
				Z:         3.0,
				Timestamp: fixedUnix,
			},
			setupMock: func(useCase *mocks.MockCreateGyroscopeUseCase) {
				gyro := &domain.Gyroscope{
					ID:        "mock-id",
					DeviceID:  "00:0a:95:9d:68:16",
					X:         1.0,
					Y:         2.0,
					Z:         3.0,
					Timestamp: fixedTimestamp,
				}
				useCase.On("Execute", mock.MatchedBy(func(dto domain.GyroscopeDto) bool {
					return dto.DeviceID == "00:0a:95:9d:68:16" &&
						dto.X == 1.0 &&
						dto.Y == 2.0 &&
						dto.Z == 3.0 &&
						dto.Timestamp == fixedUnix
				})).Return(gyro, nil).Once()
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":        "mock-id",
				"device_id": "00:0a:95:9d:68:16",
				"x":         1.0,
				"y":         2.0,
				"z":         3.0,
				"timestamp": fixedTimestamp.Format(time.RFC3339),
			},
		},

		{
			name: "Repository error",
			input: domain.GyroscopeDto{
				DeviceID:  "00:0a:95:9d:68:16",
				X:         1.0,
				Y:         2.0,
				Z:         3.0,
				Timestamp: fixedUnix,
			},
			setupMock: func(useCase *mocks.MockCreateGyroscopeUseCase) {
				useCase.On("Execute", mock.Anything).Return(nil, domain.ErrSaveGyroscopeData).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": domain.ErrSaveGyroscopeData.Error(), // "failed to save gyroscope data"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializar mock
			useCase := &mocks.MockCreateGyroscopeUseCase{}
			tt.setupMock(useCase)

			// Configurar router
			router := setupGyroscopeRouter(useCase)

			// Criar requisição
			body, _ := json.Marshal(tt.input)
			req, _ := http.NewRequest(http.MethodPost, "/api/telemetry/gyroscope", bytes.NewReader(body))
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
