package codetop

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		if len(word) == 0 {
			return true
		}
		return false
	}
	for i, line := range board {
		for j, b := range line {
			if b == word[0] {
				if findWord(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func findWord(board [][]byte, word string, i, j, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	board[i][j] = '0'
	if findWord(board, word, i-1, j, k+1) {
		return true
	}
	if findWord(board, word, i+1, j, k+1) {
		return true
	}
	if findWord(board, word, i, j-1, k+1) {
		return true
	}
	if findWord(board, word, i, j+1, k+1) {
		return true
	}
	board[i][j] = word[k]
	return false
}
