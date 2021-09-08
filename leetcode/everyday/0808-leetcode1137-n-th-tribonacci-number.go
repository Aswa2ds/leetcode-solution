package everyday

func tribonacci(n int) int {
	var trib [3]int
	trib[0], trib[1], trib[2] = 0, 1, 1
	if n < 3 {
		return trib[n]
	}

	for i := 3; i <= n; i++ {
		trib[2], trib[1], trib[0] = trib[2] + trib[1] + trib[0], trib[2], trib[1]
	}
	return trib[2]
}
