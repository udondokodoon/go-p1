package gop1_test

import (
	. "github.com/udondokodoon/gop1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gop1", func() {
	Describe("gop1", func() {
		Context("hogehoge", func() {
			It("should be 'Hello!! hogehoge'", func() {
				Expect(HelloWorld("hogehoge")).To(Equal("Hello!! hogehoge"))
			})
		})
	})
})
