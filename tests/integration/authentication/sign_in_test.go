package authentication_test

import (
	"testing"

	. "github.com/bed72/oohferta/src/data/models/responses"
	. "github.com/bed72/oohferta/src/domain/constants"
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

			value := UnmarshalMock[ResponseModel[ErrorResponseModel]](
				ResponseModel[ErrorResponseModel]{},
				response,
			)

			Expect(response.StatusCode()).To(Equal(StatusBadRequest))

			Expect(err).To(BeNil())

			Expect(value.IsSuccess).To(BeFalse())
			Expect(value.Data.Message).To(Equal("Oops! Algo deu um passeio fora dos trilhos."))
			Expect(value.Data.Description).To(Equal("Invalid login credentials"))
		})
		It("With successful return, account registered", func() {
			response, err := RequestToTest.SetBody(RegisteredAccountMock()).Post(signInURL)

			value := UnmarshalMock[ResponseModel[AuthenticationResponseModel]](
				ResponseModel[AuthenticationResponseModel]{},
				response,
			)

			Expect(response.StatusCode()).To(Equal(StatusOK))

			Expect(err).To(BeNil())

			Expect(value.IsSuccess).To(BeTrue())
			Expect(value.Data.Email).To(Equal("email@email.com"))
			Expect(value.Data.Id.Value()).To(Equal("642b155d-4020-4ee1-9bc4-abf7ddc9afbf"))
		})

		When("Should validate the body sing in return", func() {
			It("Null body", func() {
				response, err := RequestToTest.
					SetBody(nil).
					Post(signInURL)

				value := UnmarshalMock[ResponseModel[ErrorResponseModel]](
					ResponseModel[ErrorResponseModel]{},
					response,
				)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())

				Expect(value.IsSuccess).To(BeFalse())

				Expect(value.Data.Message).To(Equal("Oops! Algo deu um passeio fora dos trilhos."))
				Expect(value.Data.Description).To(Equal("unexpected end of JSON input"))
			})
			It("Empty email and password", func() {
				response, err := RequestToTest.
					SetBody(EmptyEmailAndPasswordAccountMock()).
					Post(signInURL)

				value := UnmarshalMock[ResponseModel[[]ErrorResponseModel]](
					ResponseModel[[]ErrorResponseModel]{},
					response,
				)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())

				Expect(len(value.Data)).To(Equal(2))
				Expect(value.IsSuccess).To(BeFalse())

				Expect(value.Data[0].Message).To(Equal("O campo Email não é válido."))
				Expect(value.Data[0].Description).To(Equal("A validação: [required] não foi atendida."))

				Expect(value.Data[1].Message).To(Equal("O campo Password não é válido."))
				Expect(value.Data[1].Description).To(Equal("A validação: [required] não foi atendida."))
			})
			It("Invalid email", func() {
				response, err := RequestToTest.
					SetBody(InvalidEmailAccountMock()).
					Post(signInURL)

				value := UnmarshalMock[ResponseModel[[]ErrorResponseModel]](
					ResponseModel[[]ErrorResponseModel]{},
					response,
				)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())

				Expect(len(value.Data)).To(Equal(1))
				Expect(value.IsSuccess).To(BeFalse())

				Expect(value.Data[0].Message).To(Equal("O campo Email não é válido."))
				Expect(value.Data[0].Description).To(Equal("A validação: [email] não foi atendida."))
			})
			It("Invalid password", func() {
				response, err := RequestToTest.
					SetBody(InvalidPasswordAccountMock()).
					Post(signInURL)

				value := UnmarshalMock[ResponseModel[[]ErrorResponseModel]](
					ResponseModel[[]ErrorResponseModel]{},
					response,
				)

				Expect(response.StatusCode()).To(Equal(StatusUnprocessableEntity))

				Expect(err).To(BeNil())

				Expect(len(value.Data)).To(Equal(1))
				Expect(value.IsSuccess).To(BeFalse())

				Expect(value.Data[0].Message).To(Equal("O campo Password não é válido."))
				Expect(value.Data[0].Description).To(Equal("A validação: [min] não foi atendida."))
			})
		})
	})
})
