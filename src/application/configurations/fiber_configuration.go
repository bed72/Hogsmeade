package configurations

import (
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
		AppName: "OhhFerta",
		// ErrorHandler: errorHandler,
	}
}
