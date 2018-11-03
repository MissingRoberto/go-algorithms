package dynamic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/dynamic"
)

var _ = Describe("Longestpathmatrixincreasing", func() {
	It("works", func() {
		m := [][]int{[]int{1, 2, 9}, []int{5, 3, 8}, []int{4, 6, 7}}
		Expect(FindLongestPathIncreasing(m)).To(Equal(4))
	})
})
