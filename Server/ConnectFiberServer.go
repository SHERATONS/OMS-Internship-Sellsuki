package Server

import (
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer() *FiberServer {
	app := fiber.New()
	return &FiberServer{app: app}
}

func (s *FiberServer) Start(Port string) error {
	return s.app.Listen(Port)
}
