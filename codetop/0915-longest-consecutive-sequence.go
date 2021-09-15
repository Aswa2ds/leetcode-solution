package codetop

func longestConsecutive(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] = 1
	}
	ret := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for k := range m {
		_, ok := m[k-1]
		if ok {
			continue
		}
		length := 0
		for _, ok := m[k]; ok; {
			length++
			k++
			_, ok = m[k]
		}
		ret = max(ret, length)
	}
	return ret
}
