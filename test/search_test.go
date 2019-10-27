package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AlgoCafe/boyer-moore-horspool/bmh"
)

var _ = Describe("Search", func() {
	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as argument", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabde")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(15))
		})
	})

	Context("Fail", func() {
		It("with \"abc abcdab abcdabcdabde\" as argument", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabdr")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(-1))
		})
	})
})