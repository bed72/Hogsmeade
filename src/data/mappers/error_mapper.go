package mappers

import (
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/bed72/oohferta/src/domain/entities"
)

func ErrorDefaultMapper(data error) responses.Response {
	return responses.Response{
		IsSuccess: false,
		Data: responses.ErrorResponseModel{
			Message:     "",
			Description: data.Error(),
		},
	}
}

func ErrorsMapper(data []responses.ErrorResponseModel) responses.Response {
	return responses.Response{
		IsSuccess: false,
		Data:      data,
	}
}

func ErrorMapper(data *entities.ErrorEntity) responses.Response {
	var message = ""
	var description = ""

	if data.Message != nil {
		message = *data.Message
	}

	if data.Description != nil {
		description = *data.Description
	}

	return responses.Response{
		IsSuccess: false,
		Data: responses.ErrorResponseModel{
			Message:     message,
			Description: description,
		},
	}
}
