package offer

func myPow(x float64, n int) float64 {
	var mul float64 = 1
	var flag bool
	if n == 0 {
		return mul
	}
	if n < 0 {
		flag = true
		n = -n
	}
	for i := 0; i < n; i++ {
		mul *= x
	}
	if flag {
		mul = 1 / mul
	}
	return mul
}
