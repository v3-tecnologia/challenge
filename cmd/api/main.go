package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ricardoraposo/challenge/internal/server"
)

func main() {
	app := server.New()
	app.RegisterRoutes()

	portEntry := os.Getenv("PORT")
	if portEntry == "" {
		portEntry = "8080"
	}

	app.Listen(fmt.Sprintf(":%s", portEntry))
}
