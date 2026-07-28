package valid_sudoku

func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := range 9 {
		for j := range 9 {
			num := board[i][j]
			if num == '.' {
				continue
			}

			boxIndex := (i/3)*3 + j/3

			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}

	return true
}
