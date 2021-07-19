package offer

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	s1, s2 := 1, 1
	for i := 2; i <= n; i++ {
		s1, s2 = s2, (s1+s2)%1000000007
	}
	return s2
}
