package routes

import (
	"github.com/bed72/oohferta/src/data/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoute(app fiber.Router, handler handlers.AuthenticationHandler) {
	route := app.Group("/authentication")

	route.Post("/sign_in", handler.SignIn).Name("Sign in")
}
