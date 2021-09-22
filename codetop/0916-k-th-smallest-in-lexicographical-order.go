package codetop

func findKthNumber(n int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	getCount := func(prefix int) int {
		cur, next := prefix, prefix+1
		count := 0
		for cur <= n {
			count += min(next, n+1) - cur
			cur *= 10
			next *= 10
		}
		return count
	}

	preNum := 1
	parent := 1
	for preNum < k {
		count := getCount(parent)
		if count+preNum > k {
			parent *= 10
			preNum++
		} else {
			parent++
			preNum += count
		}
	}
	return parent
}
