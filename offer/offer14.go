package offer

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	mul := 1
	for n > 4 {
		mul = (mul * 3)%1000000007
		n -= 3
	}
	return (mul * n)%1000000007
}

//func cuttingRope(n int) int {
//	if n == 2 {
//		return 1
//	}
//	if n == 3 {
//		return 2
//	}
//	var ropes []int
//	toRopes(&ropes, n)
//	mul := 1
//	for _, rope := range ropes {
//		fmt.Println(rope)
//		mul *= rope
//	}
//	return mul
//}
//
//func toRopes(ropes *[]int, n int) {
//	if n == 2 {
//		*ropes = append(*ropes, 2)
//		return
//	}
//	if n == 3 {
//		*ropes = append(*ropes, 3)
//		return
//	}
//	if n == 4 {
//		*ropes = append(*ropes, 4)
//		return
//	}
//	*ropes = append(*ropes, 3)
//	toRopes(ropes, n-3)
//	return
//}
