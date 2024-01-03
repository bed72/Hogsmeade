package handlers

import (
	"strings"

	models "github.com/bed72/oohferta/src/data/models/requests"
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
	body := &models.SignUpRequestModel{}

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	failures := h.validator.HasErrors(body)
	if failures != nil {
		return &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: strings.Join(failures, ""),
		}
	}

	success, failure, err := h.repository.SignIn(constants.SIGN_IN_URL, *body)

	if failure != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(failure)
	}

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(success)
}
