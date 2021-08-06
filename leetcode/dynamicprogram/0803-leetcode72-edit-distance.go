package dynamicprogram

import "math"

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			dp[i+1][j+1] = math.MaxInt32
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			}
			// 删除
			if dp[i][j+1] < dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i][j+1] + 1
			}
			// 插入
			if dp[i+1][j] < dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i+1][j] + 1
			}
			// 替换
			if dp[i][j] < dp[i+1][j+1] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
