package handler

import (
	"context"
	"fmt"

	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing/usecase/data"
	"github.com/gofiber/fiber/v2"
)

type BillingHandler struct {
	createBillingUseCase *usecase.CreateBillingUseCase
}

func NewBillingHandler(createBillingUseCase *usecase.CreateBillingUseCase) *BillingHandler {
	return &BillingHandler{createBillingUseCase}
}

func (h *BillingHandler) Index(c *fiber.Ctx) error {
	input := data.CreateBillingInput{}
	output, err := h.createBillingUseCase.Execute(context.Background(), input)
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	fmt.Println(output)
	return c.JSON(fiber.Map{"message": "Index"})
}

func (h *BillingHandler) Store(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Store"})
}

func (h *BillingHandler) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}

func (h *BillingHandler) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}

func (h *BillingHandler) Show(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}
