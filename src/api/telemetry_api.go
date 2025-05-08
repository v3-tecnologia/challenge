package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"v3-backend-challenge/db"
	"v3-backend-challenge/dto"
	"v3-backend-challenge/model"
	"v3-backend-challenge/repository"
	"v3-backend-challenge/utils"
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
	request, err := handleGenericPostBadRequest[dto.Gyroscope](c)
	if err != nil {
		return
	}

	gyroscope := model.Gyroscope{}
	gyroscope.AxisY = request.AxisY
	gyroscope.AxisX = request.AxisX
	gyroscope.AxisZ = request.AxisZ
	gyroscope.DateTimeCollected = request.DateTimeCollected
	gyroscope.MacAddr = request.MacAddr

	repo := repository.GyroscopeRepository{DB: db.DB}
	err = repo.Save(&gyroscope)
	if err != nil {
		log.Println("Error saving gyroscope data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar os dados"})
		return
	}
	c.JSON(201, gin.H{})
}

func handleGps(c *gin.Context) {
	request, err := handleGenericPostBadRequest[dto.GPS](c)
	if err != nil {
		return
	}

	gps := model.GPS{}
	gps.Latitude = request.Latitude
	gps.Longitude = request.Longitude
	gps.DateTimeCollected = request.DateTimeCollected
	gps.MacAddr = request.MacAddr

	repo := repository.GpsRepository{DB: db.DB}
	err = repo.Save(&gps)
	if err != nil {
		log.Println("Error saving GPS data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar os dados"})
		return
	}
	c.JSON(201, gin.H{})
}

func handlePhoto(c *gin.Context) {
	datetimeCollectedStr := c.PostForm("datetime_collected")
	macAddr := c.PostForm("mac_addr")
	if datetimeCollectedStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datetime_collected não fornecido"})
		return
	}
	datetimeCollected, err := time.Parse(time.RFC3339, datetimeCollectedStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datetime_collected inválido"})
		return
	}

	if macAddr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mac address não fornecido"})
		return
	}

	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Foto não fornecida"})
		return
	}

	photo := model.Photo{}
	photo.Name = file.Filename
	photo.DateTimeCollected = datetimeCollected
	photo.MacAddr = macAddr
	photo.Image, err = utils.FileHeaderToBytes(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler a imagem"})
		return
	}

	repo := repository.PhotoRepository{DB: db.DB}
	err = repo.Save(&photo)
	if err != nil {
		log.Println("Error saving photo data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar os dados"})
		return
	}

	c.JSON(201, gin.H{})
}

func handleGenericPostBadRequest[T any](c *gin.Context) (T, error) {
	var payload T

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Payload não está conforme o esperado"})
		return payload, err
	}

	return payload, nil
}
