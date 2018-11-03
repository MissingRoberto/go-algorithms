package dynamic

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NumPartitions(n int) float64 {
	dp := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]float64, n+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = 1
		for j := i; j <= n; j++ {
			dp[i][j] += dp[i-1][j] + dp[min(j-i, i)][j-i]
		}
	}
	return dp[n][n]
}
