package offer

import (
	"fmt"
	"strconv"
)

//normal
//func printNumbers(n int) []int {
//	ints := make([]int, int(math.Pow(10, float64(n))-1))
//	for i := range ints {
//		ints[i] = i+1
//	}
//	return ints
//}

//big int not dfs but bfsanddfs
//func printNumbers(n int) []int {
//	base := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
//	res := dfsPrintNumbers(base, base, n, 1)[1:]
//	var ints []int
//	for _, num := range res {
//		//fmt.Println(num)
//		a, _ := strconv.Atoi(num)
//		ints = append(ints, a)
//	}
//	return ints
//}
//
//func dfsPrintNumbers(base []string, external []string, n int, depth int) []string {
//	var res []string
//	if depth >= n {
//		return base
//	}
//	for _, b := range base {
//		for _, e := range external {
//			res = append(res, b+e)
//		}
//	}
//	return dfsPrintNumbers(res, external, n, depth+1)
//}

func printNumbers(n int) []int {
	num := make([]byte, n)
	base := make([]byte, 10)
	var res []int

	for i := 0; i < len(base); i++ {
		base[i] = byte('0' + i)
		fmt.Println(base[i])
	}

	dfsPrintNumbers(&res, num, base, 0, n)

	for _, re := range res {
		fmt.Println(re)
	}

	return nil
}

func dfsPrintNumbers(nums *[]int, num []byte, base []byte, depth int, n int) {
	if depth == n {
		atoi, _ := strconv.Atoi(string(num))
		*nums = append(*nums, atoi)
		return
	}
	for _, b := range base {
		num[depth] = b
		dfsPrintNumbers(nums, num, base, depth+1, n)
	}
}