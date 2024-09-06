package handler

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper         *errormapper.Mapper
	createUserUseCase   *usecase.CreateUserUseCase
	updateUserUseCase   *usecase.UpdateUserUseCase
	findUserByIDUseCase *usecase.FindUserByIDUseCase
}

func NewUserHandler(
	errorMapper *errormapper.Mapper,
	createUserUseCase *usecase.CreateUserUseCase,
	updateUserUseCase *usecase.UpdateUserUseCase,
	findUserByIDUseCase *usecase.FindUserByIDUseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		createUserUseCase,
		updateUserUseCase,
		findUserByIDUseCase,
	}
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	var (
		input  dto.CreateUserInput
		output dto.CreateUserOutput
	)

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err = h.createUserUseCase.Execute(c.Context(), input)
	if err != nil {
		rError := h.errorMapper.MapErrorToResponseError(err)
		return response.HandleErrorResponse(c, rError)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": output})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	var (
		input dto.UpdateUserInput
	)

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	input.UserID = idUser

	err = h.updateUserUseCase.Execute(c.Context(), input)
	if err != nil {
		rError := h.errorMapper.MapErrorToResponseError(err)
		return response.HandleErrorResponse(c, rError)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *UserHandler) Show(c *fiber.Ctx) error {
	var (
		input  dto.FindUserByIDInput
		output dto.FindUserByIDOutput
	)

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	input.UserID = idUser
	output, err = h.findUserByIDUseCase.Execute(c.Context(), input)
	if err != nil {
		rError := h.errorMapper.MapErrorToResponseError(err)
		return response.HandleErrorResponse(c, rError)
	}

	return c.JSON(fiber.Map{"data": output})
}
