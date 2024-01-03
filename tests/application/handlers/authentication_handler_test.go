package handlers_test

import (
	"testing"

	"github.com/bed72/oohferta/src/application/configurations"
	. "github.com/bed72/oohferta/src/domain/constants"
	. "github.com/bed72/oohferta/tests"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthenticationHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Handler Suite")
}

var _ = BeforeSuite(func() {
	configurations.EnvConfiguration()
})

var _ = Describe("Authentication Handler", func() {
	Context("Should validate the sing in return", func() {
		It("With succefful return", func() {
			// var data entities.AuthenticationEntity

			response, _ := RequestMock.SetBody(BodyMock).Post(SIGN_IN_URL)

			// if err := json.Unmarshal(response.Body(), &data); err != nil {
			// 	return
			// }

			Expect(response).To(Not(BeNil()))
		})
	})
})
