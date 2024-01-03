package test

import (
	"fmt"

	"github.com/jarcoal/httpmock"
)

func getHttpMock(name string) string {
	return httpmock.File(fmt.Sprintf("mocks/%s_response_mock.json", name)).String()
}

func ResponderMock(status int, url, name, method string) {
	responder := httpmock.NewStringResponder(status, getHttpMock(name))
	httpmock.RegisterResponder(method, url, responder)
}
