package algorithms_test

import (
	"testing"

	. "github.com/jszroberto/golang-algorithms"
)

func benchmarkNumberPartitions(b *testing.B, N int) {
	for n := 0; n < b.N; n++ {
		NumPartitions(N)
	}
}

func benchmarkNumberPartitionsBruteForce(b *testing.B, N int) {
	for n := 0; n < b.N; n++ {
		NumPartitionsBruteForce(N)
	}
}

func BenchmarkNumberPartitions_10(b *testing.B) { benchmarkNumberPartitions(b, 10) }
func BenchmarkNumberPartitions_20(b *testing.B) { benchmarkNumberPartitions(b, 20) }
func BenchmarkNumberPartitions_30(b *testing.B) { benchmarkNumberPartitions(b, 30) }
func BenchmarkNumberPartitions_40(b *testing.B) { benchmarkNumberPartitions(b, 40) }
func BenchmarkNumberPartitions_50(b *testing.B) { benchmarkNumberPartitions(b, 50) }

func BenchmarkNumberPartitionsBruteForce_10(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 10) }
func BenchmarkNumberPartitionsBruteForce_20(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 20) }
func BenchmarkNumberPartitionsBruteForce_30(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 30) }
func BenchmarkNumberPartitionsBruteForce_40(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 40) }
func BenchmarkNumberPartitionsBruteForce_50(b *testing.B) { benchmarkNumberPartitionsBruteForce(b, 50) }
