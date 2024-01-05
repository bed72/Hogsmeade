package handlers_test

import (
	"testing"

	. "github.com/bed72/oohferta/src/domain/constants"
	"github.com/bed72/oohferta/src/domain/entities"
	. "github.com/bed72/oohferta/tests"
	. "github.com/bed72/oohferta/tests/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthenticationHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Handler Suite")
}

var _ = Describe("Authentication Handler", func() {
	Context("Should validate the sing in return", func() {
		signInURL := AuthenticationURL + "/sign_in"

		It("With failure return, account not registered", func() {
			response, err := RequestToTest.SetBody(NotRegisteredAccountMock()).Post(signInURL)

			data := UnmarshalMock[entities.ErrorEntity](entities.ErrorEntity{}, response)

			Expect(response.StatusCode()).To(Equal(StatusBadRequest))

			Expect(err).To(BeNil())

			Expect(data.Message).To(BeNil())
			Expect(*data.Error).To(Equal("invalid_grant"))
			Expect(*data.Description).To(Equal("Invalid login credentials"))
		})
		It("With successful return, account registered", func() {
			response, err := RequestToTest.SetBody(RegisteredAccountMock()).Post(signInURL)

			data := UnmarshalMock[entities.AuthenticationEntity](entities.AuthenticationEntity{}, response)

			Expect(response.StatusCode()).To(Equal(StatusOK))

			Expect(err).To(BeNil())

			Expect(data.User.Email).To(Equal("email@email.com"))
			Expect(data.User.Id.Value()).To(Equal("642b155d-4020-4ee1-9bc4-abf7ddc9afbf"))
		})

		When("Should validate the body sing in return", func() {
			It("Empty email and password", func() {
				response, err := RequestToTest.
					SetBody(EmptyEmailAndPasswordAccountMock()).
					Post(signInURL)

				// data := UnmarshalMock[entities.ErrorEntity](entities.ErrorEntity{}, response)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())
				// Expect(data.Message).To(BeNil())

				// Expect(*data.Error).To(Equal("invalid_grant"))
				// Expect(*data.Description).To(Equal("Invalid login credentials"))
			})
			It("Invalid email", func() {
				response, err := RequestToTest.
					SetBody(InvalidEmailAccountMock()).
					Post(signInURL)

				// data := UnmarshalMock[entities.ErrorEntity](entities.ErrorEntity{}, response)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())
				// Expect(data.Message).To(BeNil())

				// Expect(*data.Error).To(Equal("invalid_grant"))
				// Expect(*data.Description).To(Equal("Invalid login credentials"))
			})
			It("Invalid password", func() {
				response, err := RequestToTest.
					SetBody(InvalidPasswordAccountMock()).
					Post(signInURL)

				// data := UnmarshalMock[entities.ErrorEntity](entities.ErrorEntity{}, response)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())
				// Expect(data.Message).To(BeNil())

				// Expect(*data.Error).To(Equal("invalid_grant"))
				// Expect(*data.Description).To(Equal("Invalid login credentials"))
			})
		})
	})
})
