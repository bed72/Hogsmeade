package repositories_test

import (
	"testing"

	. "github.com/bed72/oohferta/src/domain/constants"
	"github.com/bed72/oohferta/src/infrastructure/clients"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	. "github.com/bed72/oohferta/tests"
	. "github.com/bed72/oohferta/tests/mocks"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type request struct {
	key   string
	value string
}

func (r *request) Request() *resty.Request {
	return RequestToTest
}

func TestAuthenticationRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Repository Suite")
}

var _ = BeforeSuite(func() {
	httpmock.ActivateNonDefault(ClientToTest.GetClient())
})

var _ = AfterSuite(func() {
	httpmock.Deactivate()
})

var _ = Describe("Authentication Repository", func() {
	var requestToTest clients.RequestClient
	var repository repositories.AuthenticationRepository

	BeforeEach(func() {
		httpmock.Reset()

		requestToTest = &request{
			key:   "key_mock",
			value: "value_mock",
		}
		repository = repositories.New(requestToTest)
	})

	Context("Should validate the sign in return", func() {
		It("With error return", func() {
			ResponderMock(StatusBadRequest, SignInURL, "authentication_error", MethodPost)

			success, failure, err := repository.SignIn(SignInURL, RegisteredAccountMock())

			Expect(err).To(BeNil())
			Expect(success).To(BeNil())

			Expect(failure.Message).To(BeNil())
			Expect(*failure.Error).To(Equal("invalid_grant"))
			Expect(*failure.Description).To(Equal("Invalid login credentials"))
		})
		It("With unexpected error return", func() {
			ResponderMock(StatusBadRequest, SignInURL, "authentication_unexpected_error", MethodPost)

			success, failure, err := repository.SignIn(SignInURL, RegisteredAccountMock())

			Expect(err).To(BeNil())
			Expect(success).To(BeNil())

			Expect(failure.Message).To(BeNil())
			Expect(failure.Error).To(BeNil())
			Expect(failure.Description).To(BeNil())
		})
		It("With succefful return", func() {
			ResponderMock(StatusOK, SignInURL, "authentication_success", MethodPost)

			success, failure, err := repository.SignIn(SignInURL, RegisteredAccountMock())

			Expect(err).To(BeNil())
			Expect(failure).To(BeNil())

			Expect(success).To(Not(BeNil()))
			Expect(success.ExpiresIn).To(Equal(3600))
			Expect(success.AccessToken).To(Equal("eyJhbGciOi..."))
			Expect(success.RefreshToken).To(Equal("eyJhbGciOi..."))
			Expect(success.User.Email).To(Equal("email@email.com"))
			Expect(success.User.ID.Value()).To(Equal("69131ddb-9d69-4643-b352-dc56f3f68588"))
		})
	})
})
