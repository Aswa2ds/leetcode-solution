package codetop

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	x0 := float64(x+1)
	x1 := float64(x)
	for x0-x1 > 0.0001 {
		x0 = x1
		x1 = float64(x)/(x0*2) + x0/2
	}
	return int(x1)
}
