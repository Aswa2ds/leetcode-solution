package offer

//func myPow(x float64, n int) float64 {
//	var mul float64 = 1
//	var flag bool
//	if n == 0 {
//		return mul
//	}
//	if n < 0 {
//		flag = true
//		n = -n
//	}
//	for i := 0; i < n; i++ {
//		mul *= x
//	}
//	if flag {
//		mul = 1 / mul
//	}
//	return mul
//}

// quick power divide-conquer x(11) = x(1*1) * x(2*1) * x(4*0) * x(8*1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1/x
		n = -n
	}
	var mul float64 = 1
	for n > 0 {
		y := n & 1
		if y == 1 {
			mul *= x
		}
		n >>= 1
		x *= x
	}
	return mul
}