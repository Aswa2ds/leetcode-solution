package codetop

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	genSubsets(&ret, nums, make([]int, 0), -1)
	return ret
}

func genSubsets(ret *[][]int, nums []int, base []int, k int) {
	dest := make([]int, len(base))
	copy(dest, base)
	*ret = append(*ret, dest)
	j := len(base)
	base = append(base, 0)
	for i := k + 1; i < len(nums); i++ {
		base[j] = nums[i]
		genSubsets(ret, nums, base, i)
	}
}
