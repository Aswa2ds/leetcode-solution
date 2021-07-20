package offer

import "fmt"

func sum(i, j int) int {
	return i % 10 + i / 10 + j % 10 + j / 10
}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(&visited, m, n, 0, 0, k)
}

func dfs(visited *[][]bool, m int, n int, i int, j int, k int) int {
	if i >= m || j >= n {
		return 0
	}
	if (*visited)[i][j] {
		return 0
	}
	if sum(i, j) > k {
		return 0
	}
	(*visited)[i][j] = true
	fmt.Printf("(%d, %d)", i, j)
	return 1 + dfs(visited, m, n, i+1, j, k) + dfs(visited, m, n, i, j+1, k)
}
