package gop1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGop1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gop1 Suite")
}
