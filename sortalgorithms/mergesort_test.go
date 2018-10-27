package sortalgorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/sortalgorithms"
)

var _ = Describe("Mergesort", func() {
	It("can sort the array", func() {
		arr := []int{1, 4, 3, 2}
		Expect(MergeSort(arr)).To(Equal([]int{1, 2, 3, 4}))
		arr = []int{1, 4, 3, 2, 7}
		Expect(MergeSort(arr)).To(Equal([]int{1, 2, 3, 4, 7}))
		Expect(MergeSort([]int{})).To(Equal([]int{}))
		Expect(MergeSort([]int{1})).To(Equal([]int{1}))
	})
})
