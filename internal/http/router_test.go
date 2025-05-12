package http_test

import (
	"encoding/json"
	netHttp "net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/http"
	"github.com/wellmtx/challenge/internal/http/controller"
)

func TestRouter(t *testing.T) {
	// Initialize the gyroscope and geolocation controllers mocked
	gyroscopeController := &controller.GyroscopeController{}
	geolocationController := &controller.GeolocationController{}
	photoController := &controller.PhotoController{}
	router := http.NewRouter(gyroscopeController, geolocationController, photoController)

	go router.Init()

	// Add a sleep to allow the server to start
	time.Sleep(3 * time.Second)

	// Test the /ping endpoint
	resp, err := netHttp.Get("http://localhost:8080/ping")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, netHttp.StatusOK, resp.StatusCode)

	var bodyResp struct {
		Message string `json:"message"`
	}
	json.NewDecoder(resp.Body).Decode(&bodyResp)
	require.Equal(t, "pong", bodyResp.Message)
}
