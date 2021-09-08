package codetop


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int,0)
	dp = append(dp, nums[0])
	for _, num := range nums {
		if num > dp[len(dp)-1] {
			dp = append(dp, num)
		} else {
			lo, hi := 0, len(dp)
			for lo < hi {
				mid := (lo + hi) >> 1
				if dp[mid] >= num {
					hi = mid
				} else {
					lo = mid + 1
				}
			}
			dp[lo] = num
		}
	}
	return len(dp)
}
