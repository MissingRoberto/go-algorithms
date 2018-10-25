package algorithms_test


import (
	. "github.com/jszroberto/go-algorithms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RotateArray", func() {
	It("can rotate array", func() {
    arr := []int{1,2,3,4,5}
    Rotate(2, &arr)
    Expect(arr).To(Equal([]int{4,5,1,2,3}))
	})
  
	It("can rotate array when rotation is bigger than length", func() {
    arr := []int{1,2,3,4,5}
    Rotate(7, &arr)
    Expect(arr).To(Equal([]int{4,5,1,2,3}))
	})
})