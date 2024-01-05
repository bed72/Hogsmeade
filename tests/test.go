package tests_test

import (
	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/go-resty/resty/v2"
)

var (
	ClientToTest  = resty.New()
	RequestToTest = ClientToTest.R().SetDebug(false)
)

func ErrorRef(message, description string) *entities.ErrorEntity {
	return &entities.ErrorEntity{
		Error:       nil,
		Message:     &message,
		Description: &description,
	}
}
