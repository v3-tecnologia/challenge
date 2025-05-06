package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/database"
)

type FiberServer struct {
	*fiber.App
	*database.Database
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(),
        Database: database.NewDatabase(),
	}
}
