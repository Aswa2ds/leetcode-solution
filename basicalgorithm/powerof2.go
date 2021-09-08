package basicalgorithm

//func PowerOf2(a int) bool {
//	result := 0
//	for a != 0 {
//		result += a & 1
//		a >>= 1
//	}
//	return result == 1
//}

func PowerOf2(a int) bool {
	return a & (a-1) == 0
}