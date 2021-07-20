package offer

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m + 1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n + 1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2] //  s: aa p: aa.*
				if matches(i, j - 1) {	// s: ab p: a.*
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]  // s: ab p: ab
			}
		}
	}
	return f[m][n]
}

//func isMatch(s string, p string) bool {
//	str := []byte(s)
//	pattern := []byte(p)
//
//	return match(str, pattern)
//}
//
//func match(str []byte, pattern []byte) bool {
//	if len(str) == 0 && len(pattern) == 0 {
//		return true
//	}
//	if len(str) != 0 && len(pattern) == 0 {
//		return false
//	}
//	if len(pattern) == 1 {
//		if len(str) != 1 {
//			return false
//		} else {
//			return pattern[0] == '.' || pattern[0] == str[0]
//		}
//	}
//	var firstPattern []byte
//	if pattern[1] == '*' {
//		if len(str) == 0 {
//			return match(str, pattern[2:])
//		}
//		firstPattern = pattern[0:2]
//	} else {
//		firstPattern = pattern[0:1]
//		if len(str) == 0 {
//			return false
//		}
//		return match(str[0:1], firstPattern) && match(str[1:], pattern[1:])
//	}
//
//	var dotMean byte
//	maxMatchLength := 0
//	if firstPattern[0] == '.' {
//		maxMatchLength = len(str)
//	} else {
//		dotMean = firstPattern[0]
//		for _, b := range str {
//			if b != dotMean {
//				break
//			}
//			maxMatchLength++
//		}
//	}
//
//	for i := maxMatchLength; i >= 0; i-- {
//		if match(str[i:], pattern[2:]) {
//			return true
//		}
//	}
//	return false
//}


