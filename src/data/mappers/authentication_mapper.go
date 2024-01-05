package mappers

import (
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/bed72/oohferta/src/domain/entities"
)

func SignInMapper(data *entities.AuthenticationEntity) responses.Response {
	return responses.Response{
		IsSuccess: true,
		Data: responses.AuthenticationResponseModel{
			Id:           data.User.Id,
			ExpiresIn:    data.ExpiresIn,
			Email:        data.User.Email,
			AccessToken:  data.AccessToken,
			RefreshToken: data.RefreshToken,
		},
	}
}
