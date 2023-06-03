package main

func solve(board [][]byte) {
	width := len(board[0])
	height := len(board)

	for x := 0; x < width; x++ {
		dfs(board, x, 0, width, height)
		dfs(board, x, height-1, width, height)
	}
	for y := 0; y < height; y++ {
		dfs(board, 0, y, width, height)
		dfs(board, width-1, y, width, height)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch board[y][x] {
			case 'R':
				board[y][x] = 'O'
			case 'O':
				board[y][x] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x int, y int, width int, height int) {
	if x < 0 || y < 0 || x > width-1 || y > height-1 {
		return
	}
	if board[y][x] != 'O' {
		return
	}
	board[y][x] = 'R'
	dfs(board, x-1, y, width, height)
	dfs(board, x+1, y, width, height)
	dfs(board, x, y-1, width, height)
	dfs(board, x, y+1, width, height)
}
