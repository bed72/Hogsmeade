package handlers

import (
	"github.com/bed72/oohferta/src/data/mappers"
	"github.com/bed72/oohferta/src/data/models/requests"
	"github.com/bed72/oohferta/src/data/validators"
	"github.com/bed72/oohferta/src/domain/constants"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler interface {
	SignIn(ctx *fiber.Ctx) error
}

type authenticationHandler struct {
	validator  validators.Validator
	repository repositories.AuthenticationRepository
}

func New(
	validator validators.Validator,
	repository repositories.AuthenticationRepository,
) AuthenticationHandler {
	return &authenticationHandler{
		validator:  validator,
		repository: repository,
	}
}

func (h *authenticationHandler) SignIn(ctx *fiber.Ctx) error {
	body := &requests.SignUpRequestModel{}

	if err := ctx.BodyParser(body); err != nil {
		return ErrorHandler(ctx, constants.StatusBadRequest, mappers.ErrorDefaultMapper(err))
	}

	errs := h.validator.HasErrors(body)
	if errs != nil {
		return ErrorsHandler(ctx, constants.StatusUnprocessableEntity, mappers.ErrorsMapper(errs))
	}

	success, failure, err := h.repository.SignIn(constants.SIGN_IN_URL, *body)
	if failure != nil {
		return ErrorHandler(ctx, constants.StatusBadRequest, mappers.ErrorMapper(failure))
	}
	if err != nil {
		return ErrorHandler(ctx, constants.StatusBadRequest, mappers.ErrorDefaultMapper(err))
	}

	return SuccessHandler(ctx, constants.StatusOK, mappers.SignInMapper(success))
}
