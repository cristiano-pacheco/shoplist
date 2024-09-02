package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Store"})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}

func (h *UserHandler) Show(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}
