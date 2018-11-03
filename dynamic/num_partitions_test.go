package dynamic_test

import (
	. "github.com/jszroberto/go-algorithms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberPartitions", func() {
	It("can works for 1", func() {
		Expect(NumPartitions(1)).To(BeNumerically("==", 1))
	})

	It("can works for 3", func() {
		Expect(NumPartitions(3)).To(BeNumerically("==", 3))
	})

	It("can works for 6", func() {
		Expect(NumPartitions(6)).To(BeNumerically("==", 11))
	})

	It("can works for 30", func() {
		Expect(NumPartitions(30)).To(BeNumerically("==", 5604))
	})
})
