package configurations

import (
	"github.com/bed72/oohferta/src/application/handlers"
	"github.com/bed72/oohferta/src/data/validators"
	"github.com/bed72/oohferta/src/infrastructure/clients"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	"github.com/go-playground/validator/v10"
)

func validatorDIConfiguration() validators.Validator {
	return validators.New(validator.New())
}

func requestDIConfiguration() clients.RequestClient {
	return clients.New("apiKey", GetEnv("SUPABASE_KEY"))
}

func SignInDIConfiguration() handlers.AuthenticationHandler {
	repository := repositories.New(requestDIConfiguration())
	return handlers.New(validatorDIConfiguration(), repository)
}
