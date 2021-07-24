package offer

func exchange(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	s1, s2 := 0, len(nums)-1
	for s1 < s2 {
		for nums[s1]%2 != 0 && s1 < s2 {
			s1++
		}
		for nums[s2]%2 == 0 && s2 > s1 {
			s2--
		}
		if s1 == s2 {
			break
		}
		nums[s1], nums[s2] = nums[s2], nums[s1]
	}
	return nums
}
