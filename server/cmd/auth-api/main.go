package main

import (
	"fnb-qris-order/internal/auth"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	authHander := auth.NewHandler(auth.NewService(auth.NewRepository()))
	app.Post("/auth/login", authHandler.Login)
	log.Fatal(app.Listen(":3001"))
}
