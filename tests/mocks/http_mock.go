package mocks_test

import (
	"encoding/json"
	"fmt"

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

func UnmarshalMock[T interface{}](data T, response *resty.Response) T {
	if err := json.Unmarshal(response.Body(), &data); err != nil {
		panic(fmt.Errorf("fatal error when trying to serialize: %v", data))
	}

	return data
}
