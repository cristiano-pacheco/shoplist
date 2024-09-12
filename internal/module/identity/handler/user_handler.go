package handler

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/find_user_by_id_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/module/identity/usecase/update_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper         *errormapper.Mapper
	createUserUseCase   *create_user_usecase.UseCase
	updateUserUseCase   *update_user_usecase.UseCase
	findUserByIDUseCase *find_user_by_id_usecase.UseCase
}

func NewUserHandler(
	errorMapper *errormapper.Mapper,
	createUserUseCase *create_user_usecase.UseCase,
	updateUserUseCase *update_user_usecase.UseCase,
	findUserByIDUseCase *find_user_by_id_usecase.UseCase,
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
		input  create_user_usecase.Input
		output create_user_usecase.Output
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
		input update_user_usecase.Input
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
		input  find_user_by_id_usecase.Input
		output find_user_by_id_usecase.Output
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
