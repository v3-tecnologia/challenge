package main

import (
	"log"
	"v3-test/internal/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("Erro ao iniciar aplicação: %v", err)
	}
}
