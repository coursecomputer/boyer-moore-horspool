package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	bmh "github.com/coursecomputer/boyer-moore-horspool/source"
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
			var pattern = []byte("tuvzzzzzzzx")

			content := initContent(len(pattern))

			content[int('t')] = 10
			content[int('u')] = 9
			content[int('v')] = 8
			content[int('z')] = 1

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal(content))
		})

		It("with \"tuvzxzzzzzx\" as argument", func() {
			// Prepare
			var pattern = []byte("tuvzxzzzzzx")

			content := initContent(len(pattern))

			content[int('t')] = 10
			content[int('u')] = 9
			content[int('v')] = 8
			content[int('x')] = 6
			content[int('z')] = 1

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal(content))
		})

		It("with \"précé\" as argument", func() {
			// Prepare
			// var pattern = []byte("précé")	// can't use string for latin character
			var pattern = []byte{'p', 'r', 'é', 'c', 'é'}


			content := initContent(len(pattern))

			content[int('p')] = 4
			content[int('r')] = 3
			content[int('é')] = 2
			content[int('c')] = 1

			// Process
			table.Build(pattern)

			// Expect
			Expect(table.Content).To(Equal(content))
		})
	})
})