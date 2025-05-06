package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/wellmtx/challenge/internal/dtos"
)

func TestSaveData(t *testing.T) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(response)

	// Mock the request body
	data := dtos.GeolocationDataDto{
		Latitude:  37.7749,
		Longitude: -122.4194,
	}
	dataBytes, err := json.Marshal(data)
	require.NoError(t, err)

	context.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	context.Request.Method = "POST"
	context.Request.Body = io.NopCloser(bytes.NewBuffer(dataBytes))
	context.Request.Header.Set("Content-Type", "application/json")

	geolocationController.SaveData(context)
	require.Equal(t, 200, response.Code)
}
