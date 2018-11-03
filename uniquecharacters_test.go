package algorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms"
)

var _ = Describe("Uniquecharacters", func() {
	It("checks if unique", func() {
		Expect(IsUnique("ab")).To(BeTrue())
		Expect(IsUnique("abb")).To(BeFalse())
		Expect(IsUnique("")).To(BeTrue())
	})
})
