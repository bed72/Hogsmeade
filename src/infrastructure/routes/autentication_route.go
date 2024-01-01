package routes

import (
	"github.com/bed72/oohferta/src/infrastructure/configurations"
	"github.com/gofiber/fiber/v2"
)

func SignInRoute(router fiber.Router) {
	handler := configurations.SignInDIConfiguration()

	router.Post("/sign_in", handler.SignIn)
}
