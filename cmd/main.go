package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"v3/internal/domain"
	"v3/internal/infra/api"
	"v3/internal/infra/aws"
	"v3/internal/repository/gps"
	"v3/internal/repository/gyroscope"
	"v3/internal/repository/photo"
	usecase "v3/internal/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDSN() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		fmt.Printf("Using DATABASE_URL: %s\n", url)
		return url
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Missing required database environment variables: DB_HOST, DB_USER, DB_PASSWORD, DB_NAME")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)
	fmt.Printf("Using DSN from environment: %s\n", dsn)
	return dsn
}

func connectDB(dsn string) (*gorm.DB, error) {
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil {
				if err = sqlDB.Ping(); err == nil {
					fmt.Println("Database connection successful")
					return db, nil
				}
			}
		}
		fmt.Printf("Failed to connect to database: %v. Retrying in 5s (attempt %d/10)...\n", err, i+1)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after 10 retries: %w", err)
}

func main() {
	dsn := getDSN()
	db, err := connectDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	err = db.AutoMigrate(&domain.Gyroscope{}, &domain.GPS{}, &domain.Photo{})
	if err != nil {
		log.Fatalf("erro ao fazer AutoMigrate: %v", err)
	}

	awsService, err := aws.NewAWSService("telemetry-photos")
	if err != nil {
		log.Fatalf("failed to initialize AWS service: %v", err)
	}

	gyroRepo := gyroscope.NewGyroscopeRepository(db)
	gpsRepo := gps.NewGPSRepository(db)
	photoRepo := photo.NewPhotoRepository(db)

	createGyroUseCase := usecase.NewCreateGyroscopeUseCase(gyroRepo)
	createGPSUseCase := usecase.NewCreateGPSUseCase(gpsRepo)
	createPhotoUseCase := usecase.NewCreatePhotoUseCase(photoRepo, awsService)

	gyroHandlers := api.NewGyroscopeHandlers(createGyroUseCase)
	gpsHandlers := api.NewGPSHandlers(createGPSUseCase)
	photoHandlers := api.NewPhotoHandlers(createPhotoUseCase)

	router := api.SetupRouter(gyroHandlers, gpsHandlers, photoHandlers)
	router.Run(":8080")
}
