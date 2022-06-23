
func isValidSudoku(board [][]byte) bool {
	// validate row constraints
	for i := 0; i < len(board); i++ {
		if hasDuplicates(board, i, i+1, 0, len(board)) {
			return false
		}
	}

	for i := 0; i < len(board); i++ {
		if hasDuplicates(board, 0, len(board), i, i+1) {
			return false
		}
	}

	regionSize := int(math.Sqrt(float64(len(board))))
	for i := 0; i < regionSize; i++ {
		for j := 0; j < regionSize; j++ {
			if hasDuplicates(board, i*regionSize, (i+1)*regionSize,
				j*regionSize, (j+1)*regionSize) {
				return false
			}
		}
	}

	return true

}

func hasDuplicates(board [][]byte, startRow, endRow int, startCol, endCol int) bool {
	digits := map[byte]struct{}{}
	var b byte = '.'

	// row
	for i := startRow; i < endRow; i++ {
		for j := startCol; j < endCol; j++ {
			if board[i][j] != b {
				if _, ok := digits[board[i][j]]; ok {
					return true
				} else {
					digits[board[i][j]] = struct{}{}
				}
			}
		}
	}
	return false
}
