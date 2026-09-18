package main

import (
	"fnb-qris-order/internal/auth"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	authHandler := auth.NewHandler(auth.NewService(auth.NewRepository()))
	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/register", authHandler.Register)
	log.Fatal(app.Listen(":3001"))
}
