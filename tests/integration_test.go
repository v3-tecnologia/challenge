package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Kairum-Labs/should"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/gorm"

	"telemetry-api/internal/config"
	"telemetry-api/internal/database"
	"telemetry-api/internal/dtos/requests"
	"telemetry-api/internal/handlers"
	"telemetry-api/internal/middleware"
	"telemetry-api/internal/router"
	"telemetry-api/internal/services"
)

var (
	testDB     *gorm.DB
	testRouter *gin.Engine
	container  *postgres.PostgresContainer
	ctx        context.Context
	testNATS   *nats.Conn
)

func TestMain(m *testing.M) {
	// Setup
	ctx = context.Background()

	// Start PostgreSQL container
	var err error
	container, err = postgres.Run(ctx,
		"postgres:15-alpine",
		postgres.WithDatabase("test_telemetry_db"),
		postgres.WithUsername("test_user"),
		postgres.WithPassword("test_password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		log.Fatalf("Failed to start postgres container: %v", err)
	}

	// Get connection string (for potential future use)
	_, err = container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to get connection string: %v", err)
	}

	// Setup environment variables for test
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_NAME", "test_telemetry_db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("SERVER_PORT", "8080")

	// Connect to database
	cfg := &config.Config{
		DBHost:     "localhost",
		DBUser:     "test_user",
		DBPassword: "test_password",
		DBName:     "test_telemetry_db",
		DBPort:     "5432",
		ServerPort: "8080",
	}

	// Parse connection string to extract port
	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %v", err)
	}
	cfg.DBPort = mappedPort.Port()

	testDB, err = database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)
	testRouter = router.SetupRouter(testDB, testNATS, logger)

	// Run tests
	code := m.Run()

	// Cleanup
	if err := container.Terminate(ctx); err != nil {
		log.Printf("Failed to terminate container: %v", err)
	}

	os.Exit(code)
}

func setupTestData(t *testing.T) string {
	// Add small delay to avoid rate limiting between tests
	time.Sleep(100 * time.Millisecond)

	// Create a test user and get JWT token with unique email
	testUserEmail := fmt.Sprintf("test-%d@example.com", time.Now().UnixNano())
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("testuser-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "password123",
	}

	userBody, _ := json.Marshal(userReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(userBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusCreated)

	// Login to get token
	loginReq := requests.LoginRequest{
		Email:    testUserEmail,
		Password: "password123",
	}

	loginBody, _ := json.Marshal(loginReq)
	req = httptest.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	var loginResponse map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	should.BeEqual(t, err, nil)

	token, exists := loginResponse["access_token"].(string)
	should.BeTrue(t, exists, should.WithMessage("Access token should exist in response"))
	should.BeNotEmpty(t, token, should.WithMessage("Access token should not be empty"))

	return token
}

func TestTelemetryGyroscopeEndpoint(t *testing.T) {
	token := setupTestData(t)

	// Test valid gyroscope data
	gyroReq := requests.CreateGyroscopeRequest{
		DeviceID:  "device-123",
		X:         1.5,
		Y:         2.3,
		Z:         -0.8,
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(gyroReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusAccepted)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, messageExists := response["message"]
	should.BeTrue(t, messageExists, should.WithMessage("Response should contain message field"))
}

func TestTelemetryGyroscopeInvalidData(t *testing.T) {
	token := setupTestData(t)

	// Test invalid gyroscope data (missing required fields)
	invalidReq := map[string]interface{}{
		"device_id": "device-123",
		"x":         1.5,
		// Missing Y, Z, and timestamp
	}

	body, _ := json.Marshal(invalidReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusBadRequest)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorsExist := response["errors"]
	should.BeTrue(t, errorsExist, should.WithMessage("Response should contain errors field"))
}

func TestTelemetryGPSEndpoint(t *testing.T) {
	token := setupTestData(t)

	// Test valid GPS data
	gpsReq := requests.CreateGPSRequest{
		DeviceID:  "device-123",
		Latitude:  -23.5505,
		Longitude: -46.6333,
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(gpsReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusAccepted)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, messageExists := response["message"]
	should.BeTrue(t, messageExists, should.WithMessage("Response should contain message field"))
}

func TestTelemetryGPSInvalidData(t *testing.T) {
	token := setupTestData(t)

	// Test invalid GPS data (missing required fields)
	invalidReq := map[string]interface{}{
		"device_id": "device-123",
		"latitude":  -23.5505,
		// Missing longitude and timestamp
	}

	body, _ := json.Marshal(invalidReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/gps", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusBadRequest)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorsExist := response["errors"]
	should.BeTrue(t, errorsExist, should.WithMessage("Response should contain errors field"))
}

func TestTelemetryUnauthorizedAccess(t *testing.T) {
	// Test without token
	gyroReq := requests.CreateGyroscopeRequest{
		DeviceID:  "device-123",
		X:         1.5,
		Y:         2.3,
		Z:         -0.8,
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(gyroReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/gyroscope", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusUnauthorized)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorExists := response["error"]
	should.BeTrue(t, errorExists, should.WithMessage("Response should contain error field"))
}

func TestRateLimitingTelemetry(t *testing.T) {
	// Create a simple rate limiter specifically for this test
	rateLimiter := middleware.NewIPRateLimiter(rate.Every(time.Minute/2), 2) // Only 2 requests per 30 seconds

	// Create custom middleware for this test
	testRateLimitMiddleware := func(c *gin.Context) {
		ip := c.ClientIP()
		if !rateLimiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
				"code":  "RATE_LIMIT_EXCEEDED",
			})
			c.Abort()
			return
		}
		c.Next()
	}

	// Create test router with custom rate limiting
	testRateLimitRouter := gin.New()
	testRateLimitRouter.Use(testRateLimitMiddleware)

	// Setup services and handlers
	telemetryService := services.NewTelemetryService(testDB, zap.NewNop())
	authService := services.NewAuthService(testDB, zap.NewNop())
	telemetryHandler := handlers.NewTelemetryHandler(telemetryService, testNATS, zap.NewNop())
	authHandler := handlers.NewAuthHandler(authService, zap.NewNop())

	// Auth routes (no rate limiting for setup)
	authGroup := testRateLimitRouter.Group("v1/auth")
	authGroup.POST("/register", authHandler.CreateUser)
	authGroup.POST("/login", authHandler.Login)

	// Telemetry routes with rate limiting
	telemetryGroup := testRateLimitRouter.Group("v1/telemetry")
	telemetryGroup.Use(middleware.JWTAuth())
	telemetryGroup.Use(middleware.RequireRole("admin"))
	telemetryGroup.POST("/gyroscope", telemetryHandler.CreateGyroscope)

	// Get token for authenticated requests
	token := setupTestData(t)

	// Test rate limiting by making multiple requests quickly
	gyroReq := requests.CreateGyroscopeRequest{
		DeviceID:  "device-rate-limit",
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
		Timestamp: time.Now(),
	}

	successCount := 0
	rateLimitCount := 0

	// Make 5 requests quickly (should hit rate limit after 2)
	for i := 0; i < 5; i++ {
		body, _ := json.Marshal(gyroReq)
		req := httptest.NewRequest("POST", "/v1/telemetry/gyroscope", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		w := httptest.NewRecorder()
		testRateLimitRouter.ServeHTTP(w, req)

		switch w.Code {
		case http.StatusAccepted:
			successCount++
		case http.StatusTooManyRequests:
			rateLimitCount++

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			should.BeEqual(t, err, nil)

			_, errorExists := response["error"]
			should.BeTrue(t, errorExists, should.WithMessage("Rate limit response should contain error field"))

			_, codeExists := response["code"]
			should.BeTrue(t, codeExists, should.WithMessage("Rate limit response should contain code field"))
			should.BeEqual(t, response["code"], "RATE_LIMIT_EXCEEDED")
		}
	}

	// Should have some successful requests and some rate limited
	should.BeGreaterThan(t, successCount, 0, should.WithMessage("Should have some successful requests"))
	should.BeGreaterThan(t, rateLimitCount, 0, should.WithMessage("Should hit rate limit"))
}

// ========== AUTH TESTS ==========

func TestAuthRegisterSuccess(t *testing.T) {
	// Test successful user registration
	testUserEmail := fmt.Sprintf("new-user-%d@example.com", time.Now().UnixNano())
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("newuser-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "strongpassword123",
	}

	body, _ := json.Marshal(userReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusCreated)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	// Check response fields
	_, idExists := response["id"]
	should.BeTrue(t, idExists, should.WithMessage("Response should contain user ID"))

	_, usernameExists := response["username"]
	should.BeTrue(t, usernameExists, should.WithMessage("Response should contain username"))

	_, emailExists := response["email"]
	should.BeTrue(t, emailExists, should.WithMessage("Response should contain email"))

	_, rolesExists := response["roles"]
	should.BeTrue(t, rolesExists, should.WithMessage("Response should contain roles"))
}

func TestAuthRegisterDuplicateEmail(t *testing.T) {
	// First registration
	testUserEmail := fmt.Sprintf("duplicate-%d@example.com", time.Now().UnixNano())
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("user1-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "password123",
	}

	body, _ := json.Marshal(userReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	should.BeEqual(t, w.Code, http.StatusCreated)

	// Second registration with same email
	userReq2 := requests.CreateUserRequest{
		Username: fmt.Sprintf("user2-%d", time.Now().UnixNano()),
		Email:    testUserEmail, // Same email
		Password: "password456",
	}

	body2, _ := json.Marshal(userReq2)
	req2 := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	testRouter.ServeHTTP(w2, req2)

	should.BeEqual(t, w2.Code, http.StatusConflict)

	var response map[string]interface{}
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorExists := response["error"]
	should.BeTrue(t, errorExists, should.WithMessage("Response should contain error field"))
}

func TestAuthRegisterInvalidData(t *testing.T) {
	// Test invalid registration data
	invalidReq := map[string]interface{}{
		"username": "ab",            // Too short
		"email":    "invalid-email", // Invalid email format
		"password": "123",           // Too short
	}

	body, _ := json.Marshal(invalidReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusBadRequest)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorsExist := response["errors"]
	should.BeTrue(t, errorsExist, should.WithMessage("Response should contain errors field"))
}

func TestAuthLoginSuccess(t *testing.T) {
	// Create user first
	testUserEmail := fmt.Sprintf("login-test-%d@example.com", time.Now().UnixNano())
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("loginuser-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "correctpassword123",
	}

	body, _ := json.Marshal(userReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	should.BeEqual(t, w.Code, http.StatusCreated)

	// Now test login
	loginReq := requests.LoginRequest{
		Email:    testUserEmail,
		Password: "correctpassword123",
	}

	loginBody, _ := json.Marshal(loginReq)
	req2 := httptest.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	testRouter.ServeHTTP(w2, req2)

	should.BeEqual(t, w2.Code, http.StatusOK)

	var response map[string]interface{}
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	// Check response fields
	_, accessTokenExists := response["access_token"]
	should.BeTrue(t, accessTokenExists, should.WithMessage("Response should contain access_token"))

	_, refreshTokenExists := response["refresh_token"]
	should.BeTrue(t, refreshTokenExists, should.WithMessage("Response should contain refresh_token"))

	_, userExists := response["user"]
	should.BeTrue(t, userExists, should.WithMessage("Response should contain user"))
}

func TestAuthLoginInvalidCredentials(t *testing.T) {
	// Test login with non-existent user
	loginReq := requests.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "anypassword",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusUnauthorized)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorExists := response["error"]
	should.BeTrue(t, errorExists, should.WithMessage("Response should contain error field"))
}

func TestAuthLoginWrongPassword(t *testing.T) {
	// Create user first
	testUserEmail := fmt.Sprintf("wrong-pass-%d@example.com", time.Now().UnixNano())
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("wrongpassuser-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "correctpassword123",
	}

	body, _ := json.Marshal(userReq)
	req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	should.BeEqual(t, w.Code, http.StatusCreated)

	// Try login with wrong password
	loginReq := requests.LoginRequest{
		Email:    testUserEmail,
		Password: "wrongpassword",
	}

	loginBody, _ := json.Marshal(loginReq)
	req2 := httptest.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(loginBody))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	testRouter.ServeHTTP(w2, req2)

	should.BeEqual(t, w2.Code, http.StatusUnauthorized)

	var response map[string]interface{}
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorExists := response["error"]
	should.BeTrue(t, errorExists, should.WithMessage("Response should contain error field"))
}

func TestAuthRateLimiting(t *testing.T) {
	testUserEmail := fmt.Sprintf("rate-limit-auth-%d@example.com", time.Now().UnixNano())

	// Test rate limiting on auth endpoints
	userReq := requests.CreateUserRequest{
		Username: fmt.Sprintf("ratelimituser-%d", time.Now().UnixNano()),
		Email:    testUserEmail,
		Password: "password123",
	}

	successCount := 0
	rateLimitCount := 0

	// Make multiple registration attempts quickly
	for i := 0; i < 10; i++ {
		userReqCopy := userReq
		userReqCopy.Email = fmt.Sprintf("rate-limit-auth-%d-%d@example.com", time.Now().UnixNano(), i)

		body, _ := json.Marshal(userReqCopy)
		req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		switch w.Code {
		case http.StatusCreated:
			successCount++
		case http.StatusTooManyRequests:
			rateLimitCount++
		}
	}

	// Should have some successful requests and potentially some rate limited
	should.BeGreaterThan(t, successCount, 0, should.WithMessage("Should have some successful auth requests"))
}

// ========== TELEMETRY PHOTO TESTS ==========

func TestTelemetryPhotoEndpoint(t *testing.T) {
	token := setupTestData(t)

	// Test valid photo data (valid base64)
	validBase64Photo := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=="

	photoReq := requests.CreateTelemetryPhotoRequest{
		DeviceID:  "device-photo-123",
		Photo:     validBase64Photo,
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(photoReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusAccepted)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, messageExists := response["message"]
	should.BeTrue(t, messageExists, should.WithMessage("Response should contain message field"))
}

func TestTelemetryPhotoInvalidBase64(t *testing.T) {
	token := setupTestData(t)

	// Test invalid photo data (invalid base64)
	photoReq := requests.CreateTelemetryPhotoRequest{
		DeviceID:  "device-photo-456",
		Photo:     "invalid-base64-data!@#$",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(photoReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusBadRequest)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorExists := response["error"]
	should.BeTrue(t, errorExists, should.WithMessage("Response should contain error field"))
}

func TestTelemetryPhotoMissingFields(t *testing.T) {
	token := setupTestData(t)

	// Test missing required fields
	invalidReq := map[string]interface{}{
		"device_id": "device-photo-789",
		// Missing photo and timestamp
	}

	body, _ := json.Marshal(invalidReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusBadRequest)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	_, errorsExist := response["errors"]
	should.BeTrue(t, errorsExist, should.WithMessage("Response should contain errors field"))
}

func TestTelemetryPhotoUnauthorized(t *testing.T) {
	// Test without token
	photoReq := requests.CreateTelemetryPhotoRequest{
		DeviceID:  "device-123",
		Photo:     "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(photoReq)
	req := httptest.NewRequest("POST", "/v1/telemetry/photo", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusUnauthorized)
}

// Testes para endpoints GET

func TestGetGyroscopeData(t *testing.T) {
	token := setupTestData(t)

	// Primeiro, criar alguns dados de teste via worker (simulando processamento)
	service := services.NewTelemetryService(testDB, zap.NewNop())

	// Criar dados de teste diretamente no service (simulando processamento do worker)
	gyroReq1 := &requests.CreateGyroscopeRequest{
		DeviceID:  "device-test-1",
		X:         1.5,
		Y:         2.3,
		Z:         -0.8,
		Timestamp: time.Now().Add(-1 * time.Hour),
	}
	gyroReq2 := &requests.CreateGyroscopeRequest{
		DeviceID:  "device-test-1",
		X:         2.1,
		Y:         1.7,
		Z:         -1.2,
		Timestamp: time.Now().Add(-30 * time.Minute),
	}
	gyroReq3 := &requests.CreateGyroscopeRequest{
		DeviceID:  "device-test-2",
		X:         0.5,
		Y:         0.8,
		Z:         0.3,
		Timestamp: time.Now().Add(-15 * time.Minute),
	}

	err := service.CreateGyroscope(gyroReq1)
	should.BeEqual(t, err, nil)
	err = service.CreateGyroscope(gyroReq2)
	should.BeEqual(t, err, nil)
	err = service.CreateGyroscope(gyroReq3)
	should.BeEqual(t, err, nil)

	// Test 1: Get all gyroscope data (sem filtro de device)
	req := httptest.NewRequest("GET", "/v1/telemetry/gyroscope", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists := response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeGreaterThan(t, len(data), 2) // Pelo menos 3 registros

	pagination, exists := response["pagination"].(map[string]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, int(pagination["current_page"].(float64)), 1)
	should.BeEqual(t, int(pagination["per_page"].(float64)), 10)

	// Test 2: Get gyroscope data filtrado por device_id
	req = httptest.NewRequest("GET", "/v1/telemetry/gyroscope?device_id=device-test-1", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists = response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, len(data), 2) // Apenas 2 registros do device-test-1

	// Verificar se todos os registros são do device correto
	for _, item := range data {
		record := item.(map[string]interface{})
		should.BeEqual(t, record["device_id"], "device-test-1")
	}

	// Test 3: Test paginação
	req = httptest.NewRequest("GET", "/v1/telemetry/gyroscope?page=1&limit=1", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w = httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists = response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, len(data), 1) // Apenas 1 registro por página

	pagination, exists = response["pagination"].(map[string]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, int(pagination["per_page"].(float64)), 1)
}

func TestGetGPSData(t *testing.T) {
	token := setupTestData(t)

	// Criar dados de teste
	service := services.NewTelemetryService(testDB, zap.NewNop())

	gpsReq1 := &requests.CreateGPSRequest{
		DeviceID:  "device-gps-1",
		Latitude:  -23.5505,
		Longitude: -46.6333,
		Timestamp: time.Now().Add(-1 * time.Hour),
	}
	gpsReq2 := &requests.CreateGPSRequest{
		DeviceID:  "device-gps-1",
		Latitude:  -23.5515,
		Longitude: -46.6343,
		Timestamp: time.Now().Add(-30 * time.Minute),
	}

	err := service.CreateGPS(gpsReq1)
	should.BeEqual(t, err, nil)
	err = service.CreateGPS(gpsReq2)
	should.BeEqual(t, err, nil)

	// Test: Get GPS data
	req := httptest.NewRequest("GET", "/v1/telemetry/gps?device_id=device-gps-1", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists := response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, len(data), 2)

	// Verificar estrutura dos dados
	firstRecord := data[0].(map[string]interface{})
	should.BeNotEmpty(t, firstRecord["id"])
	should.BeEqual(t, firstRecord["device_id"], "device-gps-1")
	_, latExists := firstRecord["latitude"]
	should.BeTrue(t, latExists)
	_, lngExists := firstRecord["longitude"]
	should.BeTrue(t, lngExists)
	should.BeNotEmpty(t, firstRecord["timestamp"])
	should.BeNotEmpty(t, firstRecord["created_at"])
}

func TestGetPhotoData(t *testing.T) {
	token := setupTestData(t)

	// Criar dados de teste
	service := services.NewTelemetryService(testDB, zap.NewNop())

	photoReq1 := &requests.CreateTelemetryPhotoRequest{
		DeviceID:  "device-photo-1",
		Photo:     "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Timestamp: time.Now().Add(-1 * time.Hour),
	}

	err := service.CreateTelemetryPhoto(photoReq1)
	should.BeEqual(t, err, nil)

	// Test: Get photo data
	req := httptest.NewRequest("GET", "/v1/telemetry/photo?device_id=device-photo-1", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists := response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, len(data), 1)

	// Verificar estrutura dos dados
	firstRecord := data[0].(map[string]interface{})
	should.BeNotEmpty(t, firstRecord["id"])
	should.BeEqual(t, firstRecord["device_id"], "device-photo-1")
	should.BeNotEmpty(t, firstRecord["photo"])
	should.BeNotEmpty(t, firstRecord["timestamp"])
	should.BeNotEmpty(t, firstRecord["created_at"])
}

func TestGetDevices(t *testing.T) {
	token := setupTestData(t)

	// Criar dados de teste para múltiplos dispositivos
	service := services.NewTelemetryService(testDB, zap.NewNop())

	// Device 1: com dados de giroscópio e GPS
	gyroReq := &requests.CreateGyroscopeRequest{
		DeviceID:  "device-stats-1",
		X:         1.0,
		Y:         2.0,
		Z:         3.0,
		Timestamp: time.Now().Add(-1 * time.Hour),
	}
	gpsReq := &requests.CreateGPSRequest{
		DeviceID:  "device-stats-1",
		Latitude:  -23.5505,
		Longitude: -46.6333,
		Timestamp: time.Now().Add(-30 * time.Minute),
	}

	// Device 2: apenas com foto
	photoReq := &requests.CreateTelemetryPhotoRequest{
		DeviceID:  "device-stats-2",
		Photo:     "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Timestamp: time.Now().Add(-15 * time.Minute),
	}

	err := service.CreateGyroscope(gyroReq)
	should.BeEqual(t, err, nil)
	err = service.CreateGPS(gpsReq)
	should.BeEqual(t, err, nil)
	err = service.CreateTelemetryPhoto(photoReq)
	should.BeEqual(t, err, nil)

	// Test: Get devices
	req := httptest.NewRequest("GET", "/v1/telemetry/devices", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	should.BeEqual(t, w.Code, http.StatusOK)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	should.BeEqual(t, err, nil)

	data, exists := response["data"].([]interface{})
	should.BeTrue(t, exists)
	should.BeGreaterThan(t, len(data), 1) // Pelo menos 2 dispositivos

	// Verificar estrutura dos dados
	foundDevice1 := false
	foundDevice2 := false

	for _, item := range data {
		device := item.(map[string]interface{})
		deviceID := device["device_id"].(string)

		should.BeNotEmpty(t, device["last_seen"])
		_, totalExists := device["total_data_points"]
		should.BeTrue(t, totalExists)

		if deviceID == "device-stats-1" {
			foundDevice1 = true
			// Device 1 deve ter 1 gyroscope + 1 GPS = 2 total
			should.BeEqual(t, int(device["gyroscope_count"].(float64)), 1)
			should.BeEqual(t, int(device["gps_count"].(float64)), 1)
			should.BeEqual(t, int(device["photo_count"].(float64)), 0)
			should.BeEqual(t, int(device["total_data_points"].(float64)), 2)
		} else if deviceID == "device-stats-2" {
			foundDevice2 = true
			// Device 2 deve ter apenas 1 foto
			should.BeEqual(t, int(device["gyroscope_count"].(float64)), 0)
			should.BeEqual(t, int(device["gps_count"].(float64)), 0)
			should.BeEqual(t, int(device["photo_count"].(float64)), 1)
			should.BeEqual(t, int(device["total_data_points"].(float64)), 1)
		}
	}

	should.BeTrue(t, foundDevice1, should.WithMessage("Device device-stats-1 should be found"))
	should.BeTrue(t, foundDevice2, should.WithMessage("Device device-stats-2 should be found"))

	// Verificar paginação
	pagination, exists := response["pagination"].(map[string]interface{})
	should.BeTrue(t, exists)
	should.BeEqual(t, int(pagination["current_page"].(float64)), 1)
	should.BeEqual(t, int(pagination["per_page"].(float64)), 10)
}

func TestGetTelemetryUnauthorized(t *testing.T) {
	// Test todos os endpoints GET sem token
	endpoints := []string{
		"/v1/telemetry/gyroscope",
		"/v1/telemetry/gps",
		"/v1/telemetry/photo",
		"/v1/telemetry/devices",
	}

	for _, endpoint := range endpoints {
		req := httptest.NewRequest("GET", endpoint, nil)
		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		should.BeEqual(t, w.Code, http.StatusUnauthorized,
			should.WithMessage(fmt.Sprintf("Endpoint %s should require authorization", endpoint)))
	}
}

func TestGetTelemetryPaginationValidation(t *testing.T) {
	token := setupTestData(t)

	// Test parâmetros de paginação inválidos
	testCases := []struct {
		url          string
		expectedCode int
		description  string
	}{
		{"/v1/telemetry/gyroscope?page=0", http.StatusOK, "page=0 should default to 1"},
		{"/v1/telemetry/gyroscope?page=-1", http.StatusOK, "negative page should default to 1"},
		{"/v1/telemetry/gyroscope?limit=0", http.StatusOK, "limit=0 should default to 10"},
		{"/v1/telemetry/gyroscope?limit=200", http.StatusOK, "limit>100 should cap to 10"},
		{"/v1/telemetry/gyroscope?page=abc", http.StatusOK, "invalid page should default to 1"},
		{"/v1/telemetry/gyroscope?limit=xyz", http.StatusOK, "invalid limit should default to 10"},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", tc.url, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		should.BeEqual(t, w.Code, tc.expectedCode, should.WithMessage(tc.description))

		// Verificar se a resposta tem a estrutura correta
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		should.BeEqual(t, err, nil)

		_, dataExists := response["data"]
		should.BeTrue(t, dataExists, should.WithMessage("Response should have data field"))

		_, paginationExists := response["pagination"]
		should.BeTrue(t, paginationExists, should.WithMessage("Response should have pagination field"))
	}
}
