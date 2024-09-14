package response

import (
	"net/http"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/mapper/errormapper"
	"github.com/gofiber/fiber/v2"
)

func HandleErrorResponse(c *fiber.Ctx, rError errormapper.ResponseError) error {
	if rError.ErrorCode == errormapper.ValidationError {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"errors": rError.Errors})
	}

	if rError.ErrorCode == errormapper.AuthenticationError {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"errors": rError.Errors})
	}

	// todo inserir logger aqui
	serverError := errormapper.Error{
		Message: errormapper.ServerErrorMessage,
	}
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"errors": serverError})
}
