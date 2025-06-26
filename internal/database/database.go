package database

import (
	"fmt"

	"telemetry-api/internal/config"
	"telemetry-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	err = db.AutoMigrate(
		&models.Gyroscope{},
		&models.GPS{},
		&models.TelemetryPhoto{},
		&models.User{},
		&models.Role{},
		&models.RefreshToken{},
	)
	if err != nil {
		return nil, err
	}

	// Create admin role if it doesn't exist
	var adminRole models.Role
	result := db.Where("name = ?", "admin").First(&adminRole)
	if result.Error == gorm.ErrRecordNotFound {
		adminRole = models.Role{Name: "admin"}
		if err := db.Create(&adminRole).Error; err != nil {
			return nil, fmt.Errorf("failed to create admin role: %w", err)
		}
	}

	return db, nil
}
