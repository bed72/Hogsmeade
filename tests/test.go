package test

import (
	"fmt"

	"github.com/bed72/oohferta/src/application/configurations"
	models "github.com/bed72/oohferta/src/data/models/requests"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func getHttpMock(name string) string {
	return httpmock.File(fmt.Sprintf("mocks/%s_response_mock.json", name)).String()
}

func ResponderMock(status int, url, name, method string) {
	responder := httpmock.NewStringResponder(status, getHttpMock(name))
	httpmock.RegisterResponder(method, url, responder)
}

var (
	ClientMock  = resty.New().SetDebug(true)
	RequestMock = ClientMock.R().SetHeader("apiKey", configurations.GetEnv("SUPABASE_KEY"))
	BodyMock    = models.SignUpRequestModel{Email: "email@email.com", Password: "a1b2c3d4"}
)
