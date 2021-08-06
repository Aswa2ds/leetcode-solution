package mathskill

func mySqrt(x int) int {
	x1 := float64(x)
	for x1 * x1 - float64(x) > 0.1 {
		x1 = (float64(x)/x1+x1)/2
	}
	return int(x1)
}

//func mySqrt(x int) int {
//	if x == 0 || x == 1{
//		return x
//	}
//	low, high := 0, x
//	for low < high {
//		mid := (low + high) / 2
//		pow1 := mid * mid
//		if pow1 == x {
//			low, high = mid, mid
//			return mid
//		}
//		if pow1 > x {
//			high = mid
//		} else {
//			low = mid + 1
//		}
//	}
//	return low-1
//}

//func mySqrt(x int) int {
//	res := x/2
//	for true {
//		tmp := res * res
//		if tmp > x {
//			res /= 2
//			continue
//		} else if tmp < x {
//			tmp = (res+1)*(res+1)
//			if tmp > x {
//				return res
//			} else if tmp < x {
//				res += 1
//			} else {
//				return res+1
//			}
//		} else {
//			return res
//		}
//	}
//	return res
//}
