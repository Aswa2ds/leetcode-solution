package codetop

func findIsland(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	findIsland(grid, i+1, j)
	findIsland(grid, i, j+1)
	findIsland(grid, i-1, j)
	findIsland(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ret := 0
	for i, line := range grid {
		for j, b := range line {
			if b == '0' {
				continue
			}
			ret++
			findIsland(grid, i, j)
		}
	}
	return ret
}
