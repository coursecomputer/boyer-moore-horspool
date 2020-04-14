package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/coursecomputer/boyer-moore-horspool/bmh"
)

var _ = Describe("Search", func() {
	Context("Success", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abcdabde\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("abc abcdab abcdabcdabde")
			var pattern = []byte("abcdabde")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(15))
		})

		It("with \"uzzzzzzzzzzzzzx\" as buffer and \"zzzzzzzzzzzzzx\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("uzzzzzzzzzzzzzx")
			var pattern = []byte("zzzzzzzzzzzzzx")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(1))
		})

		It("with \"diiobfhewvkjeqwidiqwbqdhn wdq qdjioedwhjdeio cdhjklcd qwdqwxbj\" as buffer and \"deio\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("diiobfhewvkjeqwidiqwbqdhn wdq qdjioedwhjdeio cdhjklcd qwdqwxbj")
			var pattern = []byte("deio")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(40))
		})

		It("with \"bfhewvkjeqwidiqwbqdhn wdq qdjioedwhjdeio cdhjklcd qwdqwxbj\" as buffer and \"xbj\" as pattern", func() {
			// Prepare
			var index int
			var buffer = []byte("bfhewvkjeqwidiqwbqdhn wdq qdjioedwhjdeio cdhjklcd qwdqwxbj")
			var pattern = []byte("xbj")

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(55))
		})

		It("with \"les caractères précédents\" as buffer and \"précé\" as pattern (test french character)", func() {
			// Prepare
			var index int
			//var buffer = []byte("Donc aucune raison de comparer les caractères précédents")
			//var pattern = []byte("précé")		// can't use string for latin character
			var buffer = []byte{'l', 'e', 's', ' ', 'c', 'a', 'r', 'a', 'c', 't', 'è', 'r', 'e', 's', ' ', 'p', 'r', 'é', 'c', 'é', 'd', 'e', 'n', 't', 's'}
			var pattern = []byte{'p', 'r', 'é', 'c', 'é'}

			// Process
			index = bmh.Search(buffer, pattern)

			// Expect
			Expect(index).To(Equal(15))
		})
	})

	Context("Fail", func() {
		It("with \"abc abcdab abcdabcdabde\" as buffer and \"abcdabdr\" as pattern", func() {
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