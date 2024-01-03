package main

import (
	"log"

	"github.com/bed72/oohferta/src/application/configurations"
	"github.com/bed72/oohferta/src/application/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(configurations.FiberConfiguration())
	v1 := app.Group("/v1")

	configurations.EnvConfiguration()
	configurations.FiberMiddwaresConfiguration(app)

	routes.Routes(v1)

	err := app.Listen(configurations.GetEnv("PORT"))
	if err != nil {
		log.Panic(err)
	}
}
