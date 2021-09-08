package codetop

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for _, line := range matrix {
		for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}
	}
}
