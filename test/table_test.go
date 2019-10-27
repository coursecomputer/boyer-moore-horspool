package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AlgoCafe/boyer-moore-horspool/bmh"
)

// Additional function
func initContent(num int) (content [256]int) {
	for i := 0; i < 256; i++ {
		content[i] = num
	}

	return content
}


var _ = Describe("Table", func() {
	var table = bmh.Table{}

	Context("Success", func() {
		It("with \"tuvzzzzzzzx\" as argument", func() {
			// Prepare
			var content [256]int
			var pattern = []byte("tuvzzzzzzzx")

			content = initContent(len(pattern))

			content[int('t')] = 10
			content[int('u')] = 9
			content[int('v')] = 8
			content[int('z')] = 1

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal(content))
		})
	})
})