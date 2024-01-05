package validators

import (
	"fmt"

	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/go-playground/validator/v10"
)

type (
	Validator interface {
		HasErrors(data interface{}) []responses.ErrorResponseModel
		validate(data interface{}) []responses.ValidatorResponseModel
	}

	validators struct {
		validator *validator.Validate
	}
)

func New(validator *validator.Validate) Validator {
	return &validators{
		validator: validator,
	}
}

func (v *validators) HasErrors(data interface{}) []responses.ErrorResponseModel {
	if errs := v.validate(data); len(errs) > 0 && errs[0].Failure {
		errors := make([]responses.ErrorResponseModel, 0)

		for _, err := range errs {
			errors = append(
				errors,
				responses.ErrorResponseModel{
					Message:     fmt.Sprintf("O campo %s não é válido.", err.FailedField),
					Description: fmt.Sprintf("A validação: [%s] não foi atendida.", err.Tag),
				},
			)
		}

		return errors
	}

	return nil
}

func (v *validators) validate(data interface{}) []responses.ValidatorResponseModel {
	validators := []responses.ValidatorResponseModel{}
	failures := v.validator.Struct(data)

	if failures != nil {
		for _, err := range failures.(validator.ValidationErrors) {
			var element responses.ValidatorResponseModel
			element.Failure = true
			element.Tag = err.Tag()
			element.Value = err.Value()
			element.FailedField = err.Field()

			validators = append(validators, element)
		}
	}

	return validators
}
