package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/shoplist/internal/identity/application/usecase"
	"github.com/cristiano-pacheco/shoplist/internal/identity/domain/errs"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/dto"
	kernel_errs "github.com/cristiano-pacheco/shoplist/internal/kernel/errs"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/request"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/http/response"
	"github.com/cristiano-pacheco/shoplist/internal/kernel/otel"
)

type UserHandler struct {
	errorMapper         kernel_errs.ErrorMapper
	userCreateUseCase   usecase.UserCreateUseCase
	userUpdateUseCase   usecase.UserUpdateUseCase
	userFindUseCase     usecase.UserFindUseCase
	userActivateUseCase usecase.UserActivateUseCase
}

func NewUserHandler(
	errorMapper kernel_errs.ErrorMapper,
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
// @Success		201	{object}	response.Envelope[dto.CreateUserResponse]	"Successfully created user"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Create")
	defer span.End()

	var request dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserCreateInput{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	output, err := h.userCreateUseCase.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyInUse) {
			rError := h.errorMapper.MapCustomError(http.StatusBadRequest, err.Error())
			response.Error(w, rError)
			return
		}
		response.Error(w, err)
		return
	}

	resData := dto.CreateUserResponse{
		UserID: output.UserID,
		Name:   output.Name,
		Email:  output.Email,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusCreated, envelope, nil)
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
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Update")
	defer span.End()

	var updateUserRequest dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
		response.Error(w, err)
		return
	}

	id := request.Param(r, "id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserUpdateInput{
		UserID: idUser,
		Name:   updateUserRequest.Name,
	}

	if err := h.userUpdateUseCase.Execute(ctx, input); err != nil {
		response.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary		Find user
// @Description	Finds an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Success		200	{object}	response.Envelope[dto.FindUserResponse]	"Successfully found user"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/{id} [get]
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.FindByID")
	defer span.End()

	id := request.Param(r, "id")
	idUser, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserFindInput{UserID: idUser}
	output, err := h.userFindUseCase.Execute(ctx, input)
	if err != nil {
		response.Error(w, err)
		return
	}

	resData := dto.FindUserResponse{
		Name:  output.Name,
		Email: output.Email,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusOK, envelope, nil)
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
func (h *UserHandler) Activate(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Activate")
	defer span.End()

	var request dto.ActivateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserActivateUseCaseInput{UserID: request.UserID, Token: request.Token}
	if err := h.userActivateUseCase.Execute(ctx, input); err != nil {
		response.Error(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
