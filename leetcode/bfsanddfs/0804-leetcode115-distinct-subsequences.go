package bfsanddfs

func numDistinct(s string, t string) int {
	m,n := len(s), len(t)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return findNumDistinct(s, t, m-1, n-1, dp)
}

func findNumDistinct(s, t string, i, j int, dp [][]int) int {
	if j < 0 {
		return 1
	}
	if i < 0 {
		return 0
	}
	if dp[i][j] >= 0 {
		return dp[i][j]
	}
	res := findNumDistinct(s, t, i-1, j, dp)
	if s[i] == t[j] {
		res += findNumDistinct(s, t, i-1, j-1, dp)
	}
	dp[i][j] = res
	return res
}

// time 100% dp
//func numDistinct(s string, t string) int {
//	dp := make([]int, len(s)+1)
//	for i := range dp {
//		dp[i] = 1
//	}
//	last := 0
//	for j := len(t)-1; j >= 0; j-- {
//		last = 0
//		for i := len(s)-1; i >= 0; i-- {
//			if s[i] == t[j] {
//				dp[i+1], last = last, last + dp[i+1]
//			} else {
//				dp[i+1] = last
//			}
//		}
//		dp[0] = last
//	}
//	return last
//}

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
