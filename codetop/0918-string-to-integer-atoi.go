package codetop

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	ret := 0
	if len(s) == 0 {
		return ret
	}
	index := 0
	sign := 1
	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		index++
	}
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		num := int(s[index] - '0')
		ret = ret*10 + sign*num
		if ret > math.MaxInt32 {
			ret = math.MaxInt32
			break
		}
		if ret < math.MinInt32 {
			ret = math.MinInt32
			break
		}
		index++
	}
	return ret
}
