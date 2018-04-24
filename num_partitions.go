package algorithms

import (
	"fmt"
	"math"
)

// Possible overflow

func numPartitions(n, max int, prefix string, m [][]float64) float64 {
	if m[n][max] != 0 {
		return m[n][max]
	}

	if n == 0 {
		// fmt.Println(prefix)
		return 1
	}

	var count float64
	for i := math.Min(float64(n), float64(max)); i >= 1; i-- {
		count += numPartitions(n-int(i), int(i), fmt.Sprintf("%s %v", prefix, i), m)
	}

	m[n][max] = count
	return count
}

func NumPartitions(n int) float64 {
	grid := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		grid[i] = make([]float64, n+1)
	}
	return numPartitions(n, n, "", grid)
}

func numPartitionsBruteForce(n, max int, prefix string) float64 {
	if n == 0 {
		return 1
	}

	var count float64
	for i := math.Min(float64(n), float64(max)); i >= 1; i-- {
		count += numPartitionsBruteForce(n-int(i), int(i), fmt.Sprintf("%s %v", prefix, i))
	}

	return count
}

func NumPartitionsBruteForce(n int) float64 {
	return numPartitionsBruteForce(n, n, "")
}
