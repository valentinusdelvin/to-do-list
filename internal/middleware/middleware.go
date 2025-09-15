package middleware

import (
	jwt "todolist-go/internal/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type MiddlewareItf interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt *jwt.JWT
}

func NewMiddleware(jwt *jwt.JWT) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
