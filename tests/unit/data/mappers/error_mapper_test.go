package mappers_test

import (
	"errors"
	"testing"

	"github.com/bed72/oohferta/src/data/mappers"
	"github.com/bed72/oohferta/src/data/models/responses"
	"github.com/bed72/oohferta/src/domain/entities"
	tests_test "github.com/bed72/oohferta/tests"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositoryUtilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Error Mapper Suite")
}

var _ = Describe("Error Mapper", func() {
	Context("Should map the errors, method [ErrorDefaultMapper]", func() {
		It("With empty error", func() {
			data := mappers.ErrorDefaultMapper(errors.New(""))

			Expect(data.IsSuccess).To(BeFalse())

			Expect(data.Data.Message).To(Equal("Oops! Algo deu um passeio fora dos trilhos."))
			Expect(data.Data.Description).To(Equal("Parece que algo deu errado. Relaxa por um momento, logo estaremos de volta nos trilhos."))
		})
		It("With error", func() {
			data := mappers.ErrorDefaultMapper(errors.New("error mock"))

			Expect(data.IsSuccess).To(BeFalse())

			Expect(data.Data.Message).To(Equal("Oops! Algo deu um passeio fora dos trilhos."))
			Expect(data.Data.Description).To(Equal("error mock"))
		})
	})
	Context("Should map the errors, method [ErrorMapper]", func() {
		It("With empty error", func() {
			data := mappers.ErrorMapper(&entities.ErrorEntity{})

			Expect(data.IsSuccess).To(BeFalse())
			Expect(data.Data.Message).To(Equal("Oops! Algo deu um passeio fora dos trilhos."))
			Expect(data.Data.Description).To(Equal("Parece que algo deu errado. Relaxa por um momento, logo estaremos de volta nos trilhos."))
		})
		It("With error", func() {
			data := mappers.ErrorMapper(tests_test.ErrorRef("Message mock.", "Description mock."))

			Expect(data.IsSuccess).To(BeFalse())

			Expect(data.Data.Message).To(Equal("Message mock."))
			Expect(data.Data.Description).To(Equal("Description mock."))
		})
	})
	Context("Should map the errors, method [ErrorsMapper]", func() {
		It("With empty errors", func() {
			mock := []responses.ErrorResponseModel{}
			data := mappers.ErrorsMapper(mock)

			Expect(len(data.Data)).To(Equal(0))
			Expect(data.IsSuccess).To(BeFalse())
		})
		It("With two errors", func() {
			mock := []responses.ErrorResponseModel{
				{
					Message:     "Message mock.",
					Description: "Description mock.",
				},
				{
					Message:     "Other message mock.",
					Description: "Other description mock.",
				},
			}
			data := mappers.ErrorsMapper(mock)

			Expect(len(data.Data)).To(Equal(2))
			Expect(data.IsSuccess).To(BeFalse())

			Expect(data.Data[0].Message).To(Equal("Message mock."))
			Expect(data.Data[0].Description).To(Equal("Description mock."))

			Expect(data.Data[1].Message).To(Equal("Other message mock."))
			Expect(data.Data[1].Description).To(Equal("Other description mock."))
		})
	})
})
