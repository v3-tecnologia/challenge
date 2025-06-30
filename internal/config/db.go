package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := StringConec

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Falha ao conectar com o banco MySQL:", err)
	}
	RunMigrations(db)

	DB = db
	fmt.Println("✅ Banco de dados MySQL conectado com sucesso")
}
