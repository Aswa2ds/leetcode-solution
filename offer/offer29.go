package offer

type ori int

const (
	RIGHT ori = iota
	DOWN
	LEFT
	UP
)

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	state := RIGHT
	count := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var ret []int
	for i, j := 0, 0;; {
		if count == m*n {
			break
		}
		if !(visited[i][j]) {
			ret = append(ret, matrix[i][j])
			count++
			visited[i][j] = true
		}
		switch state {
		case RIGHT:
			if j+1 >= n || visited[i][j+1] {
				state = DOWN
			} else {
				j++
			}
		case DOWN:
			if i+1 >= m || visited[i+1][j] {
				state = LEFT
			} else {
				i++
			}
		case LEFT:
			if j-1 < 0 || visited[i][j-1] {
				state = UP
			} else {
				j--
			}
		case UP:
			if i-1 < 0 || visited[i-1][j] {
				state = RIGHT
			} else {
				i--
			}
		}
	}
	return ret
}
