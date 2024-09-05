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
	// todo inserir logger aqui
	serverError := errormapper.Error{
		Field:   "-",
		Message: errormapper.ServerErrorMessage,
	}
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"errors": serverError})
}
