package sortalgorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms/sortalgorithms"
)

var _ = Describe("Quicksort", func() {
	It("can sort the array", func() {
		arr := []int{1, 4, 3, 2}
		QuickSort(arr)
		Expect(arr).To(Equal([]int{1, 2, 3, 4}))
	})
})
