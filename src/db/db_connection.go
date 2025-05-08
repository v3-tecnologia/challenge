package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"v3-backend-challenge/model"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var gormConfig *gorm.Config
	if os.Getenv("DB_LOG") == "1" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)
		gormConfig = &gorm.Config{Logger: newLogger}
	} else {
		gormConfig = &gorm.Config{}
	}

	database, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = database.AutoMigrate(&model.Photo{}, &model.Gyroscope{}, &model.GPS{})
	if err != nil {
		log.Fatal("Erro ao executar migration:", err)
	}

	DB = database
}
