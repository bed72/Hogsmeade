package repositories_test

import (
	"testing"

	"github.com/bed72/oohferta/src/infrastructure/repositories"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositoryUtilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Utilities Suite")
}

var _ = Describe("Repository Utilities", func() {
	Context("Should verify if has successful body", func() {
		It("Has no body", func() {
			hasBody := repositories.HasSuccessfulBody(416)

			Expect(hasBody).To(BeFalse())
		})
		It("Has a body", func() {
			hasBody := repositories.HasSuccessfulBody(201)

			Expect(hasBody).To(BeTrue())
		})
	})
	Context("Should verify if has error body", func() {
		It("Has no body", func() {
			hasBody := repositories.HasErrorBody(416)

			Expect(hasBody).To(BeFalse())
		})
		It("Has a body", func() {
			hasBody := repositories.HasErrorBody(401)

			Expect(hasBody).To(BeTrue())
		})
	})
})
