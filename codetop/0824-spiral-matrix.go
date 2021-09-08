package codetop

type directionTyp uint8

const (
	right directionTyp = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}
	direction := right
	bu, br, bd, bl := -1, len(matrix[0]), len(matrix), -1
	i, j := 0, -1
	for bd-bu > 1 && br-bl > 1 {
		switch direction {
		case right:
			j++
			ret = append(ret, matrix[i][j])
			if j+1 >= br {
				bu++
				direction = down
			}
		case down:
			i++
			ret = append(ret, matrix[i][j])
			if i+1 >= bd {
				br--
				direction = left
			}
		case left:
			j--
			ret = append(ret, matrix[i][j])
			if j-1 <= bl {
				bd--
				direction = up
			}
		case up:
			i--
			ret = append(ret, matrix[i][j])
			if i-1 <= bu {
				bl++
				direction = right
			}
		}
	}
	return ret
}
