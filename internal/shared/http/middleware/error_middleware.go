package middleware

import (
	"github.com/cristiano-pacheco/shoplist/internal/shared/errs"
	"github.com/cristiano-pacheco/shoplist/internal/shared/http/response"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandlerMiddleware struct {
	errorMapper errs.ErrorMapper
}

func NewErrorHandlerMiddleware(errorMapper errs.ErrorMapper) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{errorMapper}
}

func (h *ErrorHandlerMiddleware) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		mappedError := h.errorMapper.Map(err)
		return response.Error(c, mappedError)
	}
}
