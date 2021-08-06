package bfs

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	

}

// out of time limit
//func numDistinct(s string, t string) int {
//	if len(s) >= 0 && len(t) == 0 {
//		return 1
//	}
//	if len(t) > len(s) {
//		return 0
//	}
//	res := 0
//	for i, cs := range s {
//		ct := int32(t[0])
//		if cs == ct {
//			res += numDistinct(s[i+1:], t[1:])
//		}
//	}
//	return res
//}
