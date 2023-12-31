package repositories

import (
	"encoding/json"

	models "github.com/bed72/oohferta/src/data/models/requests"
	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/bed72/oohferta/src/domain/paths"
	"github.com/go-resty/resty/v2"
)

type AuthenticationRepository interface {
	SignIn(body models.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.FailureEntity, error)
}

type authenticationRepository struct {
	request *resty.Request
}

func New(request *resty.Request) AuthenticationRepository {
	return &authenticationRepository{
		request: request,
	}
}

func (repository *authenticationRepository) SignIn(body models.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.FailureEntity, error) {
	response, err := repository.request.SetBody(body).Post(paths.SIGN_IN)

	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode() >= paths.StatusOK && response.StatusCode() <= paths.StatusCreated {
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
