package repositories

import (
	"encoding/json"

	models "github.com/bed72/oohferta/src/data/models/requests"
	"github.com/bed72/oohferta/src/domain/constants"
	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/bed72/oohferta/src/infrastructure/clients"
)

type AuthenticationRepository interface {
	SignIn(path string, body models.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.FailureEntity, error)
}

type authenticationRepository struct {
	request clients.RequestClient
}

func New(request clients.RequestClient) AuthenticationRepository {
	return &authenticationRepository{
		request: request,
	}
}

func (r *authenticationRepository) SignIn(url string, body models.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.FailureEntity, error) {
	response, err := r.request.Request().SetBody(body).Post(url)

	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode() >= constants.StatusOK && response.StatusCode() <= constants.StatusCreated {
		var data entities.AuthenticationEntity

		if err := json.Unmarshal(response.Body(), &data); err != nil {
			return nil, nil, err
		}

		return &data, nil, nil
	} else {
		var data entities.FailureEntity

		if err := json.Unmarshal(response.Body(), &data); err != nil {
			return nil, nil, err
		}

		return nil, &data, nil
	}
}
