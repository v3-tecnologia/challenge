package server

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(),
	}
}
