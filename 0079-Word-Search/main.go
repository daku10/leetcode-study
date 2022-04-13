package main

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])
	usedBoard := make([][]bool, height)
	for i := 0; i < height; i++ {
		usedBoard[i] = make([]bool, width)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y][x] == word[0] {
				if existBacktrack(board, width, height, word, 0, x, y, usedBoard) {
					return true
				}
			}
		}
	}
	return false
}

func existBacktrack(board [][]byte, width int, height int, word string, searchIndex int, x int, y int, usedBoard [][]bool) bool {
	if board[y][x] != word[searchIndex] {
		return false
	}
	if searchIndex == len(word)-1 {
		return true
	}
	usedBoard[y][x] = true
	if x+1 < width && usedBoard[y][x+1] == false {
		if existBacktrack(board, width, height, word, searchIndex+1, x+1, y, usedBoard) {
			return true
		}
	}
	if x-1 >= 0 && usedBoard[y][x-1] == false {
		if existBacktrack(board, width, height, word, searchIndex+1, x-1, y, usedBoard) {
			return true
		}
	}
	if y+1 < height && usedBoard[y+1][x] == false {
		if existBacktrack(board, width, height, word, searchIndex+1, x, y+1, usedBoard) {
			return true
		}
	}
	if y-1 >= 0 && usedBoard[y-1][x] == false {
		if existBacktrack(board, width, height, word, searchIndex+1, x, y-1, usedBoard) {
			return true
		}
	}
	usedBoard[y][x] = false
	return false
}
