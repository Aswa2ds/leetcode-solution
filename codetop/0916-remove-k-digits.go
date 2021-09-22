package codetop

import "strings"

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for _, digit := range num {
		for k > 0 && len(stack) > 0 && byte(digit) < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, byte(digit))
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}
	return ans
}
