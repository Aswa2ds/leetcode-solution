package codetop

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	comb := make([]int, 0)
	findCombination(&ret, comb, candidates, target, 0)
	return ret
}

func findCombination(ret *[][]int, comb []int, candidates []int, target, k int) {
	if target == 0 {
		cp := make([]int, len(comb))
		copy(cp, comb)
		*ret = append(*ret, cp)
		return
	}
	for i := k; i < len(candidates); i++ {
		if candidates[i] <= target {
			findCombination(ret, append(comb, candidates[i]), candidates, target-candidates[i], i)
		}
	}
}
