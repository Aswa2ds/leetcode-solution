package codetop

func backToOrigin(size, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, size)
	}
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := range dp[i] {
			dp[i][j] = dp[i-1][(j-1+size)%size] + dp[i-1][(j+1)%size]
		}
	}
	return dp[n][0]
}
