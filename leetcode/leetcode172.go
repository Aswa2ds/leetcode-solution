package leetcode

func trailingZeroes(n int) int {
	ret := 0
	for n > 0 {
		n /= 5
		ret += n
	}
	return ret
}
