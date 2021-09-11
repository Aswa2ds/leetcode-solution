package codetop

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{0, nums[0]}
	for i := 1; i < len(nums); i++ {
		dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
	}
	return dp[1]
}
