package main

import (
	"fmt"
	"log"
	"v3/internal/domain"
	"v3/internal/infra/api"
	"v3/internal/repository/gps"
	"v3/internal/repository/gyroscope"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	usecase "v3/internal/usecase"
)

func main() {
	dsn := "host=postgres user=meuusuario password=minhasenha dbname=meubanco port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	//  AutoMigrate para criar a tabela gyroscopes
	err = db.AutoMigrate(&domain.Gyroscope{}, &domain.GPS{})
	if err != nil {
		log.Fatalf("erro ao fazer AutoMigrate: %v", err)
	}

	// Inicialização dos repositórios e casos de uso
	gyroRepo := gyroscope.NewGyroscopeRepository(db)
	gpsRepo := gps.NewGPSRepository(db)

	createGyroUseCase := usecase.NewCreateGyroscopeUseCase(gyroRepo)
	createGPSUseCase := usecase.NewCreateGPSUseCase(gpsRepo)

	gyroHandlers := api.NewGyroscopeHandlers(createGyroUseCase)
	gpsHandlers := api.NewGPSHandlers(createGPSUseCase)

	router := api.SetupRouter(gyroHandlers, gpsHandlers)
	router.Run(":8080")
}
