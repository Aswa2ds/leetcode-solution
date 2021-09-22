package codetop

import "math"

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt64
	s1, s2 := 0, 0
	sum := nums[0]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for s1 < len(nums) {
		if sum < target {
			s1++
			if s1 >= len(nums) {
				break
			}
			sum += nums[s1]
		} else {
			minLen = min(s1-s2+1, minLen)
			sum -= nums[s2]
			s2++
		}
	}
	if minLen > len(nums) {
		minLen = 0
	}
	return minLen
}
