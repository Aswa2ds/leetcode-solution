// Package codetop provides coinchange
package codetop

import (
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
