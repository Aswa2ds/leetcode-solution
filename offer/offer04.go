package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)-1
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	//m, n := len(matrix)-1, len(matrix[0])-1
	for i, j := m, 0; i >= 0 && j <= n; {
		num := matrix[i][j]
		if num == target {
			return true
		}
		if num > target {
			i--
			continue
		}
		if num < target {
			j++
			continue
		}
	}
	return false
}
