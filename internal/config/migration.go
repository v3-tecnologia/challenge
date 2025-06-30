package config

import (
	"challenge-cloud/internal/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	models := []interface{}{
		&models.Gyroscope{},
	}
	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Fatalf("❌ Falha ao migrar %T: %v", m, err)
		}
	}

	fmt.Println("✅ Migrações realizadas com sucesso")
}
