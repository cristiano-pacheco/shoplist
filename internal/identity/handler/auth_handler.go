package handler

import (
	"net/http"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/dto"
	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/generate_token_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/errs"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/response"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	errorMapper          errs.ErrorMapperI
	generateTokenUseCase *generate_token_usecase.UseCase
}

func NewAuthHandler(
	errorMapper errs.ErrorMapperI,
	generateTokenUseCase *generate_token_usecase.UseCase,
) *AuthHandler {
	return &AuthHandler{errorMapper, generateTokenUseCase}
}

// Auth godoc
//
// @Summary		Generate authentication token
// @Description	Authenticates user credentials and returns an access token
// @Tags		Authentication
// @Accept		json
// @Produce		json
// @Param		request	body	dto.GenerateTokenRequest	true	"Login credentials (email and password)"
// @Success		200	{object}	response.Data{data=dto.GenerateTokenResponse}	"Successfully generated token"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/auth/token [post]
func (h *AuthHandler) GenerateToken(c *fiber.Ctx) error {
	var (
		input   generate_token_usecase.Input
		output  generate_token_usecase.Output
		request dto.GenerateTokenRequest
	)

	err := c.BodyParser(&request)
	if err != nil {
		return response.Error(c, err)
	}

	output, err = h.generateTokenUseCase.Execute(c.UserContext(), input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	return response.Success(c, http.StatusOK, output.Token)
}
