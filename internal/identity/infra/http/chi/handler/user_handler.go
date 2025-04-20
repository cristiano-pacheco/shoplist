package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/errs"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase"
	shared_errs "github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper         shared_errs.ErrorMapper
	userCreateUseCase   usecase.UserCreateUseCase
	userUpdateUseCase   usecase.UserUpdateUseCase
	userFindUseCase     usecase.UserFindUseCase
	userActivateUseCase usecase.UserActivateUseCase
}

func NewUserHandler(
	errorMapper shared_errs.ErrorMapper,
	userCreateUseCase usecase.UserCreateUseCase,
	userUpdateUseCase usecase.UserUpdateUseCase,
	userFindUseCase usecase.UserFindUseCase,
	userActivateUseCase usecase.UserActivateUseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		userCreateUseCase,
		userUpdateUseCase,
		userFindUseCase,
		userActivateUseCase,
	}
}

// @Summary		Create user
// @Description	Creates a new user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	dto.CreateUserRequest	true	"User data"
// @Success		201	{object}	response.Data[dto.CreateUserResponse]	"Successfully created user"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	ctx, span := otel.Trace().StartSpan(c.UserContext(), "UserHandler.Create")
	defer span.End()

	var request dto.CreateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	input := usecase.UserCreateUseCaseInput{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	output, err := h.userCreateUseCase.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyInUse) {
			return h.errorMapper.MapCustomError(http.StatusBadRequest, err.Error())
		}
		return err
	}

	resData := dto.CreateUserResponse{
		UserID: output.UserID,
		Name:   output.Name,
		Email:  output.Email,
	}

	res := response.NewData(resData)
	return c.Status(http.StatusCreated).JSON(res)
}

// @Summary		Update user
// @Description	Updates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Param		request	body	dto.UpdateUserRequest	true	"User data"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	ctx, span := otel.Trace().StartSpan(c.UserContext(), "UserHandler.Update")
	defer span.End()

	var request dto.UpdateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	input := usecase.UserUpdateUseCaseInput{
		UserID: idUser,
		Name:   request.Name,
	}

	if err := h.userUpdateUseCase.Execute(ctx, input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// @Summary		Find user
// @Description	Finds an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Success		200	{object}	response.Data[dto.FindUserResponse]	"Successfully found user"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [get]
func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	ctx, span := otel.Trace().StartSpan(c.UserContext(), "UserHandler.FindByID")
	defer span.End()

	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	input := usecase.UserFindUseCaseInput{UserID: idUser}
	output, err := h.userFindUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	resData := dto.FindUserResponse{
		Name:  output.Name,
		Email: output.Email,
	}

	res := response.NewData(resData)
	return c.Status(http.StatusOK).JSON(res)
}

// @Summary		Activate user
// @Description	Activates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	dto.ActivateUserRequest	true	"User data"
// @Success		204		"Successfully activated user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/activate [post]
func (h *UserHandler) Activate(c *fiber.Ctx) error {
	ctx, span := otel.Trace().StartSpan(c.UserContext(), "UserHandler.Activate")
	defer span.End()

	var request dto.ActivateUserRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	input := usecase.UserActivateUseCaseInput{UserID: request.UserID, Token: request.Token}
	if err := h.userActivateUseCase.Execute(ctx, input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
