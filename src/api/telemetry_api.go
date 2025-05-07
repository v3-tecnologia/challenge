package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v3-backend-challenge/src/dto"
)

func RegisterRoutes() {
	r := gin.Default()
	group := r.Group("/telemetry")
	group.POST("/gyroscope", handleGyroscope)
	group.POST("/gps", handleGps)
	group.POST("/photo", handlePhoto)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func handleGyroscope(c *gin.Context) {
	_ = handleGenericPostBadRequest[dto.Gyroscope](c)

	c.JSON(201, gin.H{})
}

func handleGps(c *gin.Context) {
	_ = handleGenericPostBadRequest[dto.GPS](c)

	c.JSON(201, gin.H{})
}

func handlePhoto(c *gin.Context) {
	timestamp := c.PostForm("timestamp")
	if timestamp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Timestamp não fornecido"})
		return
	}

	_, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Foto não fornecida"})
		return
	}

	c.JSON(201, gin.H{})
}

func handleGenericPostBadRequest[T any](c *gin.Context) T {
	var payload T

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Payload não está conforme o esperado"})
		return payload
	}

	return payload
}
