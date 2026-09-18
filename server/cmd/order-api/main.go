package main

import (
	"fnb-qris-order/internal/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/orders", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello, World!")
	})

	app.Listen(cfg.Port)
	log.Fatal(app.Listen(":3000"))
}
