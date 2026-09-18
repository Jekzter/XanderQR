package order

import "github.com/gofiber/fiber/v3"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	orders := h.service.GetAll()
	return c.JSON(orders)
}

func (h *Handler) Create(c fiber.Ctx) error {
	var order Order

	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	created := h.service.Create(order)
	return c.Status(fiber.StatusCreated).JSON(created)
}
