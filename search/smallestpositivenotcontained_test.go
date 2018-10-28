package search_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/search"
)

var _ = Describe("Smallestpositivenotcontained", func() {
	It("find the number", func() {
		Expect(FindSmallestPositiveNotContained([]int{1, 2, 7})).To(Equal(3))
	})
})
