package handlers_test

import (
	"errors"
	"testing"

	models "github.com/bed72/oohferta/src/data/models/requests"
	"github.com/bed72/oohferta/src/domain/entities"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type repositoryTest struct{}

var (
	err_mock     = errors.New("unknown error")
	failure_mock = entities.FailureEntity{
		Message: NullabeRefString("Credenciais inválidas."),
	}
	succes_mock = entities.AuthenticationEntity{
		ExpiresIn:    3600,
		RefreshToken: "eyJhbGciOiJIUzI1NiIs..",
		AccessToken:  "eyJhbGciOiJIUzI1NiIs...",
		User: entities.UserEntity{
			Email: "succes_mock@email.com",
			Id:    uuid.MustParse("69131ddb-9d69-4643-b352-dc56f3f68588"),
		},
	}
)

func NullabeRefString(value string) *string {
	return &value
}

func (repository *repositoryTest) SignIn(body models.SignUpRequestModel) (*entities.AuthenticationEntity, *entities.FailureEntity, error) {
	switch body.Email {
	case "succes_mock@email.com":
		return &succes_mock, nil, nil
	case "failure_mock@email.com":
		return nil, &failure_mock, nil
	default:
		return nil, nil, err_mock
	}
}

func TestAuthentication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Suite")
}

var _ = Describe("Authentication", func() {
	var mock repositoryTest

	BeforeEach(func() {
		mock = repositoryTest{}
	})

	Context("Try Sign In", func() {
		It("With succefful return", func() {
			body := models.SignUpRequestModel{Email: "succes_mock@email.com", Password: "P@sSw0rD"}

			success, _, _ := mock.SignIn(body)

			Expect(success.User.Email).To(Equal("succes_mock@email.com"))
		})
		It("With failure return", func() {
			body := models.SignUpRequestModel{Email: "failure_mock@email.com", Password: "P@sSw0rD"}

			_, failure, _ := mock.SignIn(body)

			Expect(failure.Error).To(BeNil())
			Expect(failure.Description).To(BeNil())
			Expect(*failure.Message).To(Equal("Credenciais inválidas."))
		})
		It("With error return", func() {
			body := models.SignUpRequestModel{Email: "error_mock@email.com", Password: "P@sSw0rD"}

			_, _, err := mock.SignIn(body)

			Expect(err).Error()
			Expect(err.Error()).To(Equal("unknown error"))
		})
	})
})
