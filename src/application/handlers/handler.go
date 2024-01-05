package handlers

import (
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, code int, data responses.Response) error {
	return ctx.Status(code).JSON(data)
}

func ErrorsHandler(ctx *fiber.Ctx, code int, data responses.Response) error {
	return ctx.Status(code).JSON(data)
}

func SuccessHandler(ctx *fiber.Ctx, code int, data responses.Response) error {
	return ctx.Status(code).JSON(data)
}
