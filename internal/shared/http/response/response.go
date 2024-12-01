package response

import (
	"net/http"

	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	rError, ok := err.(*errs.Error)
	if !ok {
		return err
	}

	if rError.Status == 0 {
		rError.Status = http.StatusInternalServerError
	}

	return c.Status(rError.Status).JSON(rError)
}

func Success(c *fiber.Ctx, status int, data interface{}) error {
	if status == 0 {
		status = http.StatusOK
	}

	return c.Status(status).JSON(Data{Data: data})
}
