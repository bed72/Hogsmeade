package main

import (
	"github.com/bed72/oohferta/src/data/handlers"
	"github.com/bed72/oohferta/src/data/validators"
	"github.com/bed72/oohferta/src/infrastructure/configurations"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	"github.com/bed72/oohferta/src/infrastructure/routes"
	v "github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(configurations.FiberConfiguration())
	v1 := app.Group("/v1")

	configurations.EnvConfiguration()
	configurations.FiberMiddwaresConfiguration(app)

	request := resty.New().R().SetHeader("apiKey", configurations.GetEnv("SUPABASE_KEY")).SetDoNotParseResponse(true)

	validator := validators.New(v.New())
	repository := repositories.New(request)
	handler := handlers.New(validator, repository)

	routes.AuthenticationRoute(v1, handler)

	app.Listen(configurations.GetEnv("PORT"))
}
