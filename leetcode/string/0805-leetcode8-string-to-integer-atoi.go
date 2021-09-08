package string

import "math"

func myAtoi(s string) int {
	index := 0
	for i, ch := range s {
		if ch != ' ' {
			index = i
			break
		}
	}
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		index++
	}
	s = s[index:]
	var ret int
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			break
		}
		ret *= 10
		ret += sign * int(ch - '0')
		if ret > math.MaxInt32 {
			ret = math.MaxInt32
			break
		}
		if ret < math.MinInt32 {
			ret = math.MinInt32
			break
		}
	}
	return ret
}
