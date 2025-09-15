package middleware

import "github.com/gofiber/fiber/v2"

func (m *Middleware) Authorization(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)

	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return ctx.Next()
}
