package main

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		if !isValidRow(board, i) {
			return false
		}
	}	

	for i := 0; i < 9; i++ {
		if !isValidColumn(board, i) {
			return false
		}
	}

	for i:= 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isValidSquare(board, i, j) {
				return false
			}			
		}
	}

	return true
}

func isValidRow(board [][]byte, index int) bool {
	items := make(map[byte]bool)
	row := board[index]
	for i := 0; i < 9; i++ {
		if row[i] == '.' {
			continue
		}
		if items[row[i]] {
			return false
		}
		items[row[i]] = true
	}
	return true
}

func isValidColumn(board [][]byte, index int) bool {
	items := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		item := board[i][index]
		if item == '.' {
			continue
		}
		if items[item] {
			return false
		}
		items[item] = true
	}
	return true
}

func isValidSquare(board [][]byte, x int, y int) bool {
	items := make(map[byte]bool)
	for i := x * 3; i < x * 3 + 3; i++ {
		for j := y * 3; j < y * 3 + 3; j++ {
			item := board[i][j]
			if item == '.' {
				continue
			}
			if items[item] {
				return false
			}
			items[item] = true
		}
	}
	return true
}
