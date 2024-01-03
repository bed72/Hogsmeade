package handlers_test

import (
	"testing"

	models "github.com/bed72/oohferta/src/data/models/requests"
	. "github.com/bed72/oohferta/src/domain/constants"
	"github.com/bed72/oohferta/src/infrastructure/clients"
	"github.com/bed72/oohferta/src/infrastructure/repositories"
	. "github.com/bed72/oohferta/tests"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var clientmock = resty.New()

type request struct {
	key   string
	value string
}

func (r *request) Request() *resty.Request {
	return clientmock.R()
}

func TestAuthentication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Suite")
}

var _ = BeforeSuite(func() {
	httpmock.ActivateNonDefault(clientmock.GetClient())
})

var _ = AfterSuite(func() {
	httpmock.Deactivate()
})

var _ = Describe("Authentication", func() {
	var requestmock clients.RequestClient
	var repository repositories.AuthenticationRepository

	BeforeEach(func() {
		httpmock.Reset()

		requestmock = &request{
			key:   "key_mock",
			value: "value_mock",
		}
		repository = repositories.New(requestmock)
	})

	Context("Should validate the Sign In repository return", func() {
		It("With succefful return", func() {
			ResponderMock(StatusOK, SIGN_IN_URL, "authentication_success", MethodPost)
			body := models.SignUpRequestModel{Email: "succes_mock@email.com", Password: "P@sSw0rD"}

			success, failure, err := repository.SignIn(SIGN_IN_URL, body)

			Expect(err).To(BeNil())
			Expect(failure).To(BeNil())

			Expect(success).To(Not(BeNil()))
			Expect(success.ExpiresIn).To(Equal(3600))
			Expect(success.AccessToken).To(Equal("eyJhbGciOi..."))
			Expect(success.RefreshToken).To(Equal("eyJhbGciOi..."))
			Expect(success.User.Email).To(Equal("email@email.com"))
			Expect(success.User.Id.Value()).To(Equal("69131ddb-9d69-4643-b352-dc56f3f68588"))
		})
		It("With failure return", func() {
			ResponderMock(StatusBadRequest, SIGN_IN_URL, "authentication_failure", MethodPost)
			body := models.SignUpRequestModel{Email: "failure_mock@email.com", Password: "P@sSw0rD"}

			success, failure, err := repository.SignIn(SIGN_IN_URL, body)

			Expect(err).To(BeNil())
			Expect(success).To(BeNil())

			Expect(failure.Message).To(BeNil())
			Expect(*failure.Error).To(Equal("invalid_grant"))
			Expect(*failure.Description).To(Equal("Invalid login credentials"))
		})
	})
})
