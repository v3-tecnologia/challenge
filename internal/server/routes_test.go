package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	s := New()
	s.RegisterRoutes()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	resp, err := s.App.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 OK")

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	expectedBody := `{"status":"ok"}`
	assert.JSONEq(t, expectedBody, string(body), "Response body does not match expected JSON")
}
