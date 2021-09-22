package codetop

// func dailyTemperatures(temperatures []int) []int {
// 	ret := make([]int, len(temperatures))
// 	if len(ret) == 0 {
// 		return ret
// 	}
// 	ret[len(ret)-1] = 0
// 	for i := len(ret) - 2; i >= 0; i-- {
// 		j := i + 1
// 		for j < len(ret) {
// 			if temperatures[j] > temperatures[i] {
// 				break
// 			}
// 			if ret[j] == 0 {
// 				j = len(ret)
// 			} else {
// 				j = j + ret[j]
// 			}
// 		}
// 		if j < len(ret) {
// 			ret[i] = j - i
// 		}
// 	}
// 	return ret
// }

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ret := make([]int, n)
	if n == 0 {
		return ret
	}
	s := make([]int, 0)
	for j, temperature := range temperatures {
		for len(s) != 0 && temperatures[s[len(s)-1]] < temperature {
			i := s[len(s)-1]
			s = s[:len(s)-1]
			ret[i] = j - i
		}
		s = append(s, j)
	}
	return ret
}
