package handler

import (
	"todolist-go/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(ctx *fiber.Ctx) error {
	param := dto.Register{}

	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"success": false,
			"message": "invalid request body",
		})
	}
	err = h.usecase.UserUsecase.Register(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"success": false,
			"message": err.Error(),
		})
	} else {
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error":   false,
			"success": true,
			"message": "user created",
		})
	}
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	param := dto.Login{}

	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"success": false,
			"message": "invalid request body",
		})
	}

	token, err := h.usecase.UserUsecase.Login(param)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"success": false,
			"message": err.Error(),
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"success": true,
			"message": "login successful",
			"token":   token,
		})
	}
}
