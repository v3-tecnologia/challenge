package main

import (
	"log"
	"os"

	"go-challenge/internal/messaging"
	"go-challenge/internal/models"
	"go-challenge/internal/repository"
	"go-challenge/internal/routes"
	"go-challenge/internal/services"
	"go-challenge/pkg/config"

	_ "go-challenge/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig()

	// Conectar ao banco de dados com GORM
	db := config.ConnectDB(cfg.DatabaseURL)

	// Realizar as migrations
	err := db.AutoMigrate(&models.Gyroscope{}, &models.GPS{}, &models.Photo{})
	if err != nil {
		log.Fatalf("Erro ao realizar as migrations: %v", err)
	}

	repo := &repository.TelemetryRepository{DB: db}

	service := services.NewTelemetryService(*repo)

	natsProducer, err := messaging.NewNATSProducer("nats://nats:4222")
	if err != nil {
		log.Fatalf("Erro ao conectar ao NATS: %v", err)
	}

	// Para produção, use Rekognition:
	// photoRecognition, err := services.NewAWSPhotoRecognition()
	// if err != nil {
	// 	log.Fatalf("Erro ao inicializar Rekognition: %v", err)
	// }
	// go natsProducer.StartPhotoConsumer(photoRecognition)

	//MOCK
	mockRecognition := &services.MockPhotoRecognition{}
	go natsProducer.StartPhotoConsumer(mockRecognition)

	r := routes.SetupRouter(service, natsProducer)

	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080"
	}
	ginSwaggerURL := ginSwagger.URL("http://" + swaggerHost + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwaggerURL))

	log.Println("Iniciando servidor na porta:", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
