package handler

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/activate_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/create_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/find_user"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/usecase/update_user"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	errorMapper         errs.ErrorMapper
	createUserUseCase   *create_user.CreateUserUseCase
	updateUserUseCase   *update_user.UpdateUserUseCase
	findUserUseCase     *find_user.FindUserUseCase
	activateUserUseCase *activate_user.ActivateUserUseCase
}

func NewUserHandler(
	errorMapper errs.ErrorMapper,
	createUserUseCase *create_user.CreateUserUseCase,
	updateUserUseCase *update_user.UpdateUserUseCase,
	findUserUseCase *find_user.FindUserUseCase,
	activateUserUseCase *activate_user.ActivateUserUseCase,
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
// @Param		request	body	dto.CreateUserRequest	true	"User data"
// @Success		201	{object}	response.Data[dto.CreateUserResponse]	"Successfully created user"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var request dto.CreateUserRequest
	err := c.BodyParser(&request)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	input := create_user.Input{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	output, err := h.createUserUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
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
	var request dto.UpdateUserRequest
	err := c.BodyParser(&request)
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

	input := update_user.Input{
		UserID: idUser,
		Name:   request.Name,
	}

	err = h.updateUserUseCase.Execute(c.UserContext(), input)
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
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Success		200	{object}	response.Data[dto.FindUserResponse]	"Successfully found user"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [get]
func (h *UserHandler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	input := find_user.Input{UserID: idUser}
	output, err := h.findUserUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
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
	var request dto.ActivateUserRequest
	err := c.BodyParser(&request)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	input := activate_user.Input{UserID: request.UserID, Token: request.Token}
	err = h.activateUserUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return c.SendStatus(http.StatusNoContent)
}
