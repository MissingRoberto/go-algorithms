package algorithms_test

import (
	"testing"

	. "github.com/jszroberto/golang-algorithms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberPartitions", func() {
	It("can works for 1", func() {
		Expect(NumPartitionsFormula(1)).To(BeNumerically("==", 1))
		Expect(NumPartitions(1)).To(BeNumerically("==", 1))
		Expect(NumPartitionsBruteForce(1)).To(BeNumerically("==", 1))
	})

	It("can works for 3", func() {
		Expect(NumPartitionsFormula(3)).To(BeNumerically("==", 3))
		Expect(NumPartitions(3)).To(BeNumerically("==", 3))
		Expect(NumPartitions(3)).To(BeNumerically("==", 3))
	})

	It("can works for 6", func() {
		Expect(NumPartitionsFormula(6)).To(BeNumerically("==", 11))
		Expect(NumPartitions(6)).To(BeNumerically("==", 11))
		Expect(NumPartitionsBruteForce(6)).To(BeNumerically("==", 11))
	})

	It("can works for 30", func() {
		Expect(NumPartitionsFormula(30)).To(BeNumerically("==", 5604))
		Expect(NumPartitions(30)).To(BeNumerically("==", 5604))
		Expect(NumPartitionsBruteForce(30)).To(BeNumerically("==", 5604))
	})
})

func benchmarkNumberPartitions(b *testing.B, N int) {
	for n := 0; n < b.N; n++ {
		NumPartitions(N)
	}
}
func benchmarkNumberPartitionsFormula(b *testing.B, N int) {
	for n := 0; n < b.N; n++ {
		NumPartitionsFormula(N)
	}
}

func benchmarkNumberPartitionsBruteForce(b *testing.B, N int) {
	for n := 0; n < b.N; n++ {
		NumPartitionsBruteForce(N)
	}
}

func BenchmarkNumberPartitions_1(b *testing.B)  { benchmarkNumberPartitions(b, 1) }
func BenchmarkNumberPartitions_10(b *testing.B) { benchmarkNumberPartitions(b, 10) }
func BenchmarkNumberPartitions_20(b *testing.B) { benchmarkNumberPartitions(b, 20) }
func BenchmarkNumberPartitions_30(b *testing.B) { benchmarkNumberPartitions(b, 30) }
func BenchmarkNumberPartitions_40(b *testing.B) { benchmarkNumberPartitions(b, 40) }
func BenchmarkNumberPartitions_50(b *testing.B) { benchmarkNumberPartitions(b, 50) }
func BenchmarkNumberPartitions_60(b *testing.B) { benchmarkNumberPartitions(b, 60) }

func BenchmarkNumberPartitionsFormula_1(b *testing.B) { benchmarkNumberPartitionsFormula(b, 1) }

func BenchmarkNumberPartitionsFormula_10(b *testing.B) { benchmarkNumberPartitionsFormula(b, 10) }

func BenchmarkNumberPartitionsFormula_20(b *testing.B) { benchmarkNumberPartitionsFormula(b, 20) }
func BenchmarkNumberPartitionsFormula_30(b *testing.B) { benchmarkNumberPartitionsFormula(b, 30) }
func BenchmarkNumberPartitionsFormula_40(b *testing.B) { benchmarkNumberPartitionsFormula(b, 40) }
func BenchmarkNumberPartitionsFormula_50(b *testing.B) { benchmarkNumberPartitionsFormula(b, 50) }
func BenchmarkNumberPartitionsFormula_60(b *testing.B) { benchmarkNumberPartitionsFormula(b, 60) }

func BenchmarkNumberPartitionsBruteForce_10(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 10) }
func BenchmarkNumberPartitionsBruteForce_20(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 20) }
func BenchmarkNumberPartitionsBruteForce_30(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 30) }
func BenchmarkNumberPartitionsBruteForce_40(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 40) }
func BenchmarkNumberPartitionsBruteForce_50(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 50) }
func BenchmarkNumberPartitionsBruteForce_60(b *testing.B) {
	benchmarkNumberPartitionsBruteForce(b, 60)
}
