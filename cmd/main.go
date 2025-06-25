package main

import (
	"log"
	"v3-test/internal/config/databases"
	"v3-test/internal/config/routers"
)

func main() {
	databases.ConnectMongo()

	server := routers.SetupRouter()

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
