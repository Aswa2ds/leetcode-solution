package codetop

import "math"

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) == 0 {
		return ret
	}
	if len(nums) == 1 {
		ret = append(ret, []int{nums[0]})
		return ret
	}
	subPermute(&ret, nums, make([]int, 0), 0)
	return ret
}

func subPermute(ret *[][]int, nums []int, res []int, depth int) {
	if depth == len(nums) {
		ans := make([]int, depth)
		copy(ans, res)
		*ret = append(*ret, ans)
		return
	}
	for i, num := range nums {
		if num != math.MinInt64 {
			nums[i] = math.MinInt64
			res = append(res, num)
			subPermute(ret, nums, res, depth+1)
			res = res[:len(res)-1]
			nums[i] = num
		}
	}
}
