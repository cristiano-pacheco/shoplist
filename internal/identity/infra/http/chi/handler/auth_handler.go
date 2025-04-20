package handler

import (
	"net/http"

	"github.com/cristiano-pacheco/shoplist/internal/identity/application/usecase"
	"github.com/cristiano-pacheco/shoplist/internal/identity/infra/http/dto"
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/cristiano-pacheco/shoplist/internal/shared/otel"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	errorMapper          errs.ErrorMapper
	tokenGenerateUseCase usecase.TokenGenerateUseCase
}

func NewAuthHandler(
	errorMapper errs.ErrorMapper,
	tokenGenerateUseCase usecase.TokenGenerateUseCase,
) *AuthHandler {
	return &AuthHandler{errorMapper, tokenGenerateUseCase}
}

// @Summary		Generate authentication token
// @Description	Authenticates user credentials and returns an access token
// @Tags		Authentication
// @Accept		json
// @Produce		json
// @Param		request	body	dto.GenerateTokenRequest	true	"Login credentials (email and password)"
// @Success		200	{object}	response.Data[dto.GenerateTokenResponse]	"Successfully generated token"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/auth/token [post]
func (h *AuthHandler) GenerateToken(c *fiber.Ctx) error {
	ctx, span := otel.Trace().StartSpan(c.UserContext(), "AuthHandler.GenerateToken")
	defer span.End()

	var request dto.GenerateTokenRequest
	err := c.BodyParser(&request)
	if err != nil {
		return response.Error(c, err)
	}

	input := usecase.TokenGenerateInput{
		Email:    request.Email,
		Password: request.Password,
	}

	output, err := h.tokenGenerateUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		return response.Error(c, rError)
	}

	res := response.NewData(dto.GenerateTokenResponse{Token: output.Token})
	return c.Status(http.StatusOK).JSON(res)
}
