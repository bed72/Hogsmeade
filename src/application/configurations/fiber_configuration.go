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

/*
func errorHandler(ctx *fiber.Ctx, err error) error {
	code := constants.StatusInternalServerError
	data := "Ops! Parece que nosso servidor está fazendo uma pausa para um cafezinho! Pedimos desculpas pelo transtorno momentâneo. 💻😊"

	var erro *fiber.Error
	if errors.As(err, &erro) {
		code = erro.Code
		data = erro.Message
	}

	err = errorType(ctx, code, data)
	if err != nil {
		return errorDefault(ctx, code, data)
	}

	return nil
}

func errorType(ctx *fiber.Ctx, code int, data string) error {
	if code == constants.StatusUnprocessableEntity {
		return errorDefault(ctx, code, strings.Split(data, "&"))
	}

	return errorDefault(ctx, code, data)
}

func errorDefault[T string | []string](ctx *fiber.Ctx, code int, message T) error {
	return ctx.Status(code).JSON(responses.Response[T]{
		IsSuccess: false,
		Data:      message,
	})
}
*/
