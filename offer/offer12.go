package offer

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
				if findWord(board, word, i, j, 0) {
					return true
				}
		}
	}
	return false
}

func findWord(board [][]byte, word string, row int, col int, index int) bool {
	if board[row][col] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	tmp := board[row][col]
	board[row][col] = ' '
	if row > 0 && findWord(board, word, row-1, col, index+1) {
		return true
	}
	if col > 0 && findWord(board, word, row, col-1, index+1) {
		return true
	}
	if row < len(board)-1 && findWord(board, word, row+1, col, index+1) {
		return true
	}
	if col < len(board[0])-1 && findWord(board, word, row, col+1, index+1) {
		return true
	}
	board[row][col] = tmp
	return false
}
