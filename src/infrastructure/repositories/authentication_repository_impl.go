package repositories

import (
	"encoding/json"

	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/go-resty/resty/v2"
)

type AuthenticationRepository interface {
	SignIn(email string, password string) (*entities.AuthenticationEntity, *entities.FailureEntity, error)
}

type authenticationRepository struct {
	request *resty.Request
}

func New(request *resty.Request) AuthenticationRepository {
	return &authenticationRepository{
		request: request,
	}
}

func (repository *authenticationRepository) SignIn(email string, password string) (*entities.AuthenticationEntity, *entities.FailureEntity, error) {
	response, err := repository.
		request.
		SetBody(map[string]interface{}{"email": email, "password": password}).
		Post("https://naiwfnrkbtsmzllsvmxj.supabase.co/auth/v1/token?grant_type=password")

	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode() >= 200 && response.StatusCode() <= 201 {
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
