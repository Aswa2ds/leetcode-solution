package offer

import (
	"fmt"
)

func available(m, n, k int) bool {
	return (m/10 + m%10 + n/10 + n%10) <= k
}

func movingCount(m int, n int, k int) int {
	//visited := make([][]bool, m)
	//for i := range visited {
	//	visited[i] = make([]bool, n)
	//}
	//return moveCount(visited, m, n, k, 0, 0)
	c := 0
	if n-1 < k {
		c = n
	}
	c = k+1
	sum := 0
	for i := 0; i < c; i++ {
		count := 0
		if k-i >= m-1 {
			count = m
		} else {
			count = k-i+1
		}
		sum += count
	}
	return sum
}

func max(args ...int) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("no input args")
	}
	m := args[0]
	for _, arg := range args {
		if arg > m {
			m = arg
		}
	}
	return m, nil
}

func moveCount(visited [][]bool, m int, n int, k int, i int, j int) int {
	if !available(i, j, k) {
		return 0
	}
	var down, right int
	visited[i][j] = true
	//if i > 0 && !visited[i-1][j] {
	//	up = moveCount(visited, m, n, k, i-1, j)
	//}
	//if j > 0 && !visited[i][j-1] {
	//	left = moveCount(visited, m, n, k, i, j-1)
	//}
	if i < m-1 && !visited[i+1][j] {
		down = moveCount(visited, m, n, k, i+1, j)
	}
	if j < n-1 && !visited[i][j+1]{
		right = moveCount(visited, m, n, k, i, j+1)
	}
	return 1 + down + right
}
