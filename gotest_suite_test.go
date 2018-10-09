package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotest Suite")
}
