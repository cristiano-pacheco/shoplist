package handler

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/activate_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/create_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/find_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/update_user_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/response"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper         errs.ErrorMapperI
	createUserUseCase   *create_user_usecase.UseCase
	updateUserUseCase   *update_user_usecase.UseCase
	findUserUseCase     *find_user_usecase.UseCase
	activateUserUseCase *activate_user_usecase.UseCase
}

func NewUserHandler(
	errorMapper errs.ErrorMapperI,
	createUserUseCase *create_user_usecase.UseCase,
	updateUserUseCase *update_user_usecase.UseCase,
	findUserUseCase *find_user_usecase.UseCase,
	activateUserUseCase *activate_user_usecase.UseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		createUserUseCase,
		updateUserUseCase,
		findUserUseCase,
		activateUserUseCase,
	}
}

// @Summary		Create user
// @Description	Creates a new user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	create_user_usecase.Input	true	"User data"
// @Success		201	{object}	response.Data{data=create_user_usecase.Output}	"Successfully created user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users [post]
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

// @Summary		Update user
// @Description	Updates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id		path	integer		true	"User ID"
// @Param		request	body	update_user_usecase.Input	true	"User data"
// @Success		204	{object}	response.Data	"Successfully updated user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [put]
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

// @Summary		Find user
// @Description	Finds an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		id		path	integer		true	"User ID"
// @Success		200	{object}	response.Data{data=find_user_usecase.Output}	"Successfully found user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [get]
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

// @Summary		Activate user
// @Description	Activates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	activate_user_usecase.Input	true	"User data"
// @Success		204		"Successfully activated user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/activate [post]
func (h *UserHandler) Activate(c *fiber.Ctx) error {
	type requestInput struct {
		UserID uint64 `json:"user_id"`
		Token  string `json:"token"`
	}
	var rInput requestInput

	err := c.BodyParser(&rInput)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	input := activate_user_usecase.Input{UserID: rInput.UserID, Token: rInput.Token}

	err = h.activateUserUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return c.SendStatus(http.StatusNoContent)
}
