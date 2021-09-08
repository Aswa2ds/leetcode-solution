package simulation

func gameOfLife(board [][]int)  {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])
	if m == 0 {
		return
	}
	calculate := func(i, j int) int {
		sum := 0
		if i > 0 {
			sum += board[i-1][j] & 1
			if j > 0 {
				sum += board[i-1][j-1] & 1
			}
			if j < m-1 {
				sum += board[i-1][j+1] & 1
			}
		}
		if i < n-1 {
			sum += board[i+1][j] & 1
			if j > 0 {
				sum += board[i+1][j-1] & 1
			}
			if j < m-1 {
				sum += board[i+1][j+1] & 1
			}
		}
		if j > 0 {
			sum += board[i][j-1] & 1
		}
		if j < m-1 {
			sum += board[i][j+1] & 1
		}
		return sum
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := calculate(i, j)
			if board[i][j] & 1 == 1 {
				if count > 3 || count < 2 {
					board[i][j] |= 0
				} else {
					board[i][j] |= 2
				}
			} else {
				if count == 3 {
					board[i][j] |= 2
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			board[i][j] >>= 1
		}
	}
}
