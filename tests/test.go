package test_test

import (
	"github.com/go-resty/resty/v2"
)

var (
	ClientToTest  = resty.New()
	RequestToTest = ClientToTest.R().SetDebug(false)
)
