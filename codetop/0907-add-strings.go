package codetop

import "strconv"

func addStrings(num1 string, num2 string) string {
	var c int
	var ret string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		result := x + y + c
		ret = strconv.Itoa(result%10) + ret
		c = result / 10
	}
	return ret
}
