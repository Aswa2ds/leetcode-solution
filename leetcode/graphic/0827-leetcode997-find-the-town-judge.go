package graphic

func findJudge(n int, trust [][]int) int {
	in, out := make([]int, n+1), make([]int, n+1)
	for _, v := range trust {
		out[v[0]]++
		in[v[1]]++
	}
	ret := -1
	for i := 1; i < n+1; i++ {
		if out[i] == 0 && in[i] == n-1 {
			if ret == -1 {
				ret = i
			} else {
				return -1
			}
		}
	}
	return ret
}
