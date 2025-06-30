package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Exemplo: usuário root, senha 123, banco telemetry, rodando local na porta padrão
	dsn := "root:@tcp(127.0.0.1:3306)/telemetry?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Falha ao conectar com o banco MySQL:", err)
	}
	RunMigrations(db)

	DB = db
	fmt.Println("✅ Banco de dados MySQL conectado com sucesso")
}
