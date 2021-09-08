package basicalgorithm

func EuclideanAlgorithm(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for {
		a, b = b, a%b
		if b == 0 {
			return a
		}
	}
}

func GengXiangJianSun(a, b int) int {
	for {
		if a == b {
			return a
		}
		if b > a {
			a, b = b, a
		}
		a, b = b, a-b
	}
}

func MaxCommonFactor(a, b int) int {
	ret := 1
	for {
		if a == b {
			break
		}
		aEven := a & 1 == 0
		bEven := b & 1 == 0
		if aEven && bEven {
			a >>= 1
			b >>= 1
			ret *= 2
		} else if aEven {
			a >>=1
		} else if bEven {
			b >>= 1
		} else {
			if b > a {
				a, b = b, a
			}
			a, b = b, a-b
		}
	}
	return ret
}