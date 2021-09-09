package codetop

import "math"

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func maxSubArray(nums []int) int {
	sum := math.MinInt32
	maxSum := math.MinInt32
	for _, num := range nums {
		sum = max(sum+num, num)
		maxSum = max(sum, maxSum)
	}
	return maxSum
}
