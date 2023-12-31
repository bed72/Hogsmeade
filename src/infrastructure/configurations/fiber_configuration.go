package configurations

import (
	"strings"

	"github.com/bed72/oohferta/src/data/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddwaresConfiguration(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
	)
}

func FiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: errorHandler,
	}
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	messages := strings.Split(err.Error(), ". ")

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(validators.GlobalFailureHandlerResponse{
		Success:  false,
		Messages: messages,
	})
}
