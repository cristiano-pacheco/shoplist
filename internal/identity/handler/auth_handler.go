package handler

import (
	"net/http"

	"github.com/cristiano-pacheco/go-modulith/internal/identity/usecase/generate_token_usecase"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/http/response"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/telemetry"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	errorMapper          *errormapper.Mapper
	generateTokenUseCase *generate_token_usecase.UseCase
}

func NewAuthHandler(
	errorMapper *errormapper.Mapper,
	generateTokenUseCase *generate_token_usecase.UseCase,
) *AuthHandler {
	return &AuthHandler{errorMapper, generateTokenUseCase}
}

func (h *AuthHandler) Execute(c *fiber.Ctx) error {
	var (
		input  generate_token_usecase.Input
		output generate_token_usecase.Output
	)

	t := telemetry.Get()
	ctx, span := t.StartSpan(c.Context(), "auth_handler.execute")
	defer span.End()

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	output, err = h.generateTokenUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.MapErrorToResponseError(err)
		return response.HandleErrorResponse(c, rError)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": output.Token})
}
