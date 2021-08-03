package leetcode

func capitalToNumber(capital byte) int {
	return int(capital - 'A' + 1)
}

func titleToNumber(columnTitle string) int {
	dec := 0
	multiple := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		capital := columnTitle[i]
		dec += capitalToNumber(capital) * multiple
		multiple = multiple<<4 + multiple<<3 + multiple<<1
	}
	return dec
}
