package repositories

import (
	"encoding/json"

	"github.com/bed72/oohferta/src/data/models/requests"
	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/bed72/oohferta/src/infrastructure/clients"
)

type AuthenticationRepository interface {
	SignIn(path string, body requests.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.ErrorEntity, error)
}

type authenticationRepository struct {
	request clients.RequestClient
}

func New(request clients.RequestClient) AuthenticationRepository {
	return &authenticationRepository{
		request: request,
	}
}

func (r *authenticationRepository) SignIn(
	url string,
	body requests.SignUpRequestModel,
) (*entities.AuthenticationEntity, *entities.ErrorEntity, error) {
	response, err := r.request.Request().SetBody(body).Post(url)
	if err != nil {
		return nil, nil, err
	}

	if HasSuccessfulBody(response.StatusCode()) {
		var data entities.AuthenticationEntity
		if err := json.Unmarshal(response.Body(), &data); err != nil {
			return nil, nil, err
		}

		return &data, nil, nil
	}

	if HasErrorBody(response.StatusCode()) {
		var data entities.ErrorEntity
		if err := json.Unmarshal(response.Body(), &data); err != nil {
			return nil, nil, err
		}

		return nil, &data, nil
	}

	return nil, &entities.ErrorEntity{}, nil
}
