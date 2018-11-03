package dynamic

import "fmt"

// https://www.geeksforgeeks.org/find-the-longest-path-in-a-matrix-with-given-constraints/

func FindLongestPathIncreasing(m [][]int) int {
	dp := make([][]int, len(m)+1)
	for i := 0; i <= len(m); i++ {
		dp[i] = make([]int, len(m[0])+1)
	}

	var max int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			sum := 1
			if i > 0 && (m[i][j] == m[i-1][j]+1 || m[i][j] == m[i-1][j]-1) {
				sum += dp[i][j+1]
			}

			if j > 0 && (m[i][j] == m[i][j-1]+1 || m[i][j] == m[i][j-1]-1) {
				sum += dp[i+1][j]
			}

			dp[i+1][j+1] = sum
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(dp)
	return max
}
