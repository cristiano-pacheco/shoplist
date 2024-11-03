package handler

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/find_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/update_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/response"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper       errs.ErrorMapperI
	createUserUseCase *create_user_usecase.UseCase
	updateUserUseCase *update_user_usecase.UseCase
	findUserUseCase   *find_user_usecase.UseCase
}

func NewUserHandler(
	errorMapper errs.ErrorMapperI,
	createUserUseCase *create_user_usecase.UseCase,
	updateUserUseCase *update_user_usecase.UseCase,
	findUserUseCase *find_user_usecase.UseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		createUserUseCase,
		updateUserUseCase,
		findUserUseCase,
	}
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	var (
		input  create_user_usecase.Input
		output create_user_usecase.Output
	)
	t := telemetry.Get()
	ctx, span := t.StartSpan(c.Context(), "user_handler.store")
	defer span.End()

	err := c.BodyParser(&input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	output, err = h.createUserUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return response.Success(c, http.StatusCreated, output)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	var (
		input update_user_usecase.Input
	)

	t := telemetry.Get()
	ctx, span := t.StartSpan(c.Context(), "user_handler.update")
	defer span.End()

	err := c.BodyParser(&input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}
	input.UserID = idUser

	err = h.updateUserUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *UserHandler) Show(c *fiber.Ctx) error {
	var (
		input  find_user_usecase.Input
		output find_user_usecase.Output
	)

	t := telemetry.Get()
	ctx, span := t.StartSpan(c.Context(), "user_handler.show")
	defer span.End()

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	input.UserID = idUser
	output, err = h.findUserUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return response.Success(c, http.StatusOK, output)
}
