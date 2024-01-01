package configurations

import (
	"github.com/bed72/oohferta/src/data/handlers"
	"github.com/bed72/oohferta/src/data/validators"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

func validatorDIConfiguration() validators.Validator {
	return validators.New(validator.New())
}

func requestDIConfiguration() *resty.Request {
	return resty.New().R().SetHeader("apiKey", GetEnv("SUPABASE_KEY"))
}

func SignInDIConfiguration() handlers.AuthenticationHandler {
	repository := repositories.New(requestDIConfiguration())
	return handlers.New(validatorDIConfiguration(), repository)
}
