package middleware

import (
	jwt "todolist-go/internal/pkg/jwt"
	"todolist-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type MiddlewareItf interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt     *jwt.JWT
	usecase usecase.Usecase
}

func NewMiddleware(jwt *jwt.JWT, usecase usecase.Usecase) *Middleware {
	return &Middleware{
		jwt:     jwt,
		usecase: usecase,
	}
}
