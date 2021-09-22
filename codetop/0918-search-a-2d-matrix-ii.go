package codetop

func searchMatrix(matrix [][]int, target int) bool {
	ret := false
	m, n := len(matrix)-1, len(matrix[0])-1
	for i, j := m, 0; i >= 0 && j <= n; {
		if matrix[i][j] == target {
			ret = true
			break
		}
		if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return ret
}
