package server

import (
	"github.com/gofiber/fiber/v2"
	"sabir222/go-backend/internal/database"
)

type FiberServer struct {
	*fiber.App
	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-backend-project",
			AppName:      "sabir-go-project",
		}),

		db: database.New(),
	}
	return server
}
