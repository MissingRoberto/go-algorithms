package algorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms"
)

var _ = Describe("Pairsum", func() {
	It("works", func() {
		arr := []int{1, 2, 3}

		i, j := FindPairWithSort(arr, 3)
		Expect(i).To(Equal(0))
		Expect(j).To(Equal(1))

		i, j = FindPairWithSort(arr, 6)
		Expect(i).To(Equal(-1))
		Expect(j).To(Equal(-1))
	})

	It("works", func() {
		arr := []int{1, 2, 3}

		i, j := FindPairHash(arr, 3)
		Expect(i).To(Equal(0))
		Expect(j).To(Equal(1))

		i, j = FindPairHash(arr, 6)
		Expect(i).To(Equal(-1))
		Expect(j).To(Equal(-1))
	})
})
