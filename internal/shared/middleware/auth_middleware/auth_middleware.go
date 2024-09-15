package auth_middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Execute(c *fiber.Ctx) error {
	fmt.Println("Auth Middleware")
	return c.Next()
}
