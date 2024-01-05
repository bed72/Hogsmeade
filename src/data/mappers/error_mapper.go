package mappers

import (
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/bed72/oohferta/src/domain/entities"
)

var (
	isSuccess = false
)

func ErrorDefaultMapper(data error) responses.ResponseModel[responses.ErrorResponseModel] {
	message := "Oops! Algo deu um passeio fora dos trilhos."
	description := "Parece que algo deu errado. Relaxa por um momento, logo estaremos de volta nos trilhos."

	if len(data.Error()) > 4 {
		description = data.Error()
	}

	return responses.ResponseModel[responses.ErrorResponseModel]{
		IsSuccess: isSuccess,
		Data: responses.ErrorResponseModel{
			Message:     message,
			Description: description,
		},
	}
}

func ErrorsMapper(data []responses.ErrorResponseModel) responses.ResponseModel[[]responses.ErrorResponseModel] {
	return responses.ResponseModel[[]responses.ErrorResponseModel]{
		IsSuccess: isSuccess,
		Data:      data,
	}
}

func ErrorMapper(data *entities.ErrorEntity) responses.ResponseModel[responses.ErrorResponseModel] {
	message := "Oops! Algo deu um passeio fora dos trilhos."
	description := "Parece que algo deu errado. Relaxa por um momento, logo estaremos de volta nos trilhos."

	if data.Message != nil {
		message = *data.Message
	}

	if data.Description != nil {
		description = *data.Description
	}

	return responses.ResponseModel[responses.ErrorResponseModel]{
		IsSuccess: isSuccess,
		Data: responses.ErrorResponseModel{
			Message:     message,
			Description: description,
		},
	}
}
