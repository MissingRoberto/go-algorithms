package search_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/search"
)

var _ = Describe("Findinrotatedarray", func() {
	It("works", func() {
		arr := []int{3, 4, 5, 6, 7, 1, 2}
		Expect(FindInRotatedArray(3, arr)).To(Equal(0))
		Expect(FindInRotatedArray(4, arr)).To(Equal(1))
		Expect(FindInRotatedArray(5, arr)).To(Equal(2))
		Expect(FindInRotatedArray(6, arr)).To(Equal(3))
		Expect(FindInRotatedArray(7, arr)).To(Equal(4))
		Expect(FindInRotatedArray(1, arr)).To(Equal(5))
		Expect(FindInRotatedArray(2, arr)).To(Equal(6))
	})
})
