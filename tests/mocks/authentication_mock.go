package mocks_test

import (
	"github.com/bed72/oohferta/src/application/configurations"
	models "github.com/bed72/oohferta/src/data/models/requests"
)

var (
	InvalidFakePasswordMock = "safe"
	InvalidFakeEmailMock    = "exempleexemple.com"
	ValidFakeEmailMock      = "exemple@exemple.com"
	ValidFakePasswordMock   = "safe_password_ooh_ferta"

	AuthenticationURL = "http://localhost:3000/v1/authentication"
)

func RegisteredAccountMock() models.SignUpRequestModel {
	configurations.EnvConfiguration()

	return models.SignUpRequestModel{
		Email:    configurations.GetEnv("SUPABASE_TEST_EMAIL"),
		Password: configurations.GetEnv("SUPABASE_TEST_PASSWORD"),
	}
}

func NotRegisteredAccountMock() models.SignUpRequestModel {
	return models.SignUpRequestModel{
		Email:    ValidFakeEmailMock,
		Password: ValidFakePasswordMock,
	}
}

func EmptyEmailAndPasswordAccountMock() models.SignUpRequestModel {
	return models.SignUpRequestModel{Email: "", Password: ""}
}

func InvalidEmailAndPasswordAccountMock() models.SignUpRequestModel {
	return models.SignUpRequestModel{Email: InvalidFakeEmailMock, Password: InvalidFakePasswordMock}
}

func InvalidPasswordAccountMock() models.SignUpRequestModel {
	return models.SignUpRequestModel{Email: ValidFakeEmailMock, Password: InvalidFakePasswordMock}
}

func InvalidEmailAccountMock() models.SignUpRequestModel {
	return models.SignUpRequestModel{Email: InvalidFakeEmailMock, Password: ValidFakePasswordMock}
}

/*

func RegisteredAccountMock() map[string]string {
	configurations.EnvConfiguration()

	return map[string]string{
		"email":    configurations.GetEnv("SUPABASE_TEST_EMAIL"),
		"password": configurations.GetEnv("SUPABASE_TEST_PASSWORD"),
	}
}

func NotRegisteredAccountMock() map[string]string {
	return map[string]string{
		"email":    ValidFakeEmailMock,
		"password": ValidFakePasswordMock,
	}
}

func EmptyEmailAndPasswordAccountMock() map[string]string {
	return map[string]string{"email": "", "password": ""}
}

func InvalidEmailAndPasswordAccountMock() map[string]string {
	return map[string]string{"email": InvalidFakeEmailMock, "password": InvalidFakePasswordMock}
}

func InvalidPasswordAccountMock() map[string]string {
	return map[string]string{"email": ValidFakeEmailMock, "password": InvalidFakePasswordMock}
}

func InvalidEmailAccountMock() map[string]string {
	return map[string]string{"email": InvalidFakeEmailMock, "password": ValidFakePasswordMock}
}

*/
