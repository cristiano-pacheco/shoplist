package handler

import (
	"net/http"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper       *errormapper.Mapper
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(
	errorMapper *errormapper.Mapper,
	createUserUseCase *usecase.CreateUserUseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		createUserUseCase,
	}
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	var input dto.CreateUserInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	output, err := h.createUserUseCase.Execute(c.Context(), input)
	if err != nil {
		rError := h.errorMapper.MapErrorToResponseError(err)
		return response.HandleErrorResponse(c, rError)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"user": output})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}

func (h *UserHandler) Show(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update"})
}
