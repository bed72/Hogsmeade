package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	Validator interface {
		HasErrors(data interface{}) []string
		validate(data interface{}) []FailureResponse
	}

	validators struct {
		validator *validator.Validate
	}

	FailureResponse struct {
		Failure     bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	GlobalFailureHandlerResponse struct {
		Success  bool     `json:"success"`
		Messages []string `json:"messages"`
	}
)

func New(validator *validator.Validate) Validator {
	return &validators{
		validator: validator,
	}
}

func (v *validators) HasErrors(data interface{}) []string {
	if errs := v.validate(data); len(errs) > 0 && errs[0].Failure {
		messages := make([]string, 0)

		for _, err := range errs {
			messages = append(messages, fmt.Sprintf(
				"Preencha um(a) %s válido(a).",
				err.FailedField,
			))
		}

		return messages

	}

	return nil
}

func (v *validators) validate(data interface{}) []FailureResponse {
	validators := []FailureResponse{}
	failures := v.validator.Struct(data)

	if failures != nil {
		for _, err := range failures.(validator.ValidationErrors) {
			var element FailureResponse

			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Value()
			element.Failure = true

			validators = append(validators, element)
		}
	}

	return validators
}
