package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Route("/authentication", SignInRoute, "Sign In")
}
