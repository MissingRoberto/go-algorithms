package algorithms_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegerPartition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegerPartition Suite")
}
