package handlers

import (
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/gofiber/fiber/v2"
)

type errorHandlerType interface {
	responses.ErrorResponseModel | []responses.ErrorResponseModel
}

func ErrorHandler[T errorHandlerType](ctx *fiber.Ctx, code int, data responses.ResponseModel[T]) error {
	return ctx.Status(code).JSON(data)
}

func SuccessHandler[T interface{}](ctx *fiber.Ctx, code int, data responses.ResponseModel[T]) error {
	return ctx.Status(code).JSON(data)
}
