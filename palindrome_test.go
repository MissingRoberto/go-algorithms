package algorithms_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jszroberto/go-algorithms"
)

var _ = Describe("Palindrome", func() {
  FIt("checks correct palindromes", func(){
    Expect(CheckPalindrome("ov")).To(BeFalse())    
    Expect(CheckPalindrome("ovo")).To(BeTrue())    
    Expect(CheckPalindrome("abba")).To(BeTrue())    
    Expect(CheckPalindrome("abcba")).To(BeTrue())    
  })
})
