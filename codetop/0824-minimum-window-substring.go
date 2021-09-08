package codetop

import "math"

func minWindow(s string, t string) string {
	need := make(map[int32]int)
	needCount := 0
	for _, ch := range t {
		need[ch]++
		needCount++
	}
	i := 0
	left, right := math.MinInt32, math.MaxInt32
	for j, ch := range s {
		if need[ch] > 0 {
			needCount--
			need[ch]--
		}
		if needCount == 0 {
			var c int32
			for {
				c = int32(s[i])
				if _, ok := need[c]; ok {
					break
				}
				i++
			}
			if j-i < right - left {
				right, left = j, i
			}
			need[c]++
			needCount++
			i++
		}
	}
	if left < 0 {
		return ""
	} else {
		return s[left:right+1]
	}
}
