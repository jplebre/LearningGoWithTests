package calculator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator", func() {
	It("Should Pass", func() {
		Expect(true).To(Equal(true))
	})
})
