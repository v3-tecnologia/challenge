package seeder

import (
	"challenge-cloud/internal/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	username := "admin"
	password := "123456"

	// Verifica se o usuário já existe
	var existing models.User
	if err := db.Where("username = ?", username).First(&existing).Error; err == nil {
		log.Printf("⚠️ Usuário admin já existe. Seed ignorado.")
		return
	}

	// Gera hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("❌ Erro ao gerar hash da senha: %v", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("❌ Erro ao criar usuário: %v", err)
	}

	fmt.Println("✅ Usuário admin criado com sucesso!")
}
