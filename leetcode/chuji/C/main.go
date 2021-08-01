package main

func main() {
	isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
}

func isValidSudoku(board [][]byte) bool {
	var h [9]map[byte]int
	var c [9]map[byte]int
	var b [9]map[byte]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if h[i] == nil {
				h[i] = make(map[byte]int)
			}
			if _, ok := h[i][board[i][j]]; ok {
				return false
			} else {
				h[i][board[i][j]] = 1
			}

			if c[j] == nil {
				c[j] = make(map[byte]int)
			}
			if _, ok := c[j][board[i][j]]; ok {
				return false
			} else {
				c[j][board[i][j]] = 1
			}
			bi := i/3 + j/3*3
			if b[bi] == nil {
				b[bi] = make(map[byte]int)
			}
			if _, ok := b[bi][board[i][j]]; ok {
				return false
			} else {
				b[bi][board[i][j]] = 1
			}
		}
	}
	return true

}
