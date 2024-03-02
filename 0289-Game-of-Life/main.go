package main

func gameOfLife(board [][]int) {
	height, width := len(board), len(board[0])
	oldBoard := make([][]int, height)
	for i := 0; i < height; i++ {
		oldBoard[i] = make([]int, width)
		copy(oldBoard[i], board[i])
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			board[y][x] = calc(oldBoard, x, y, width, height)
		}
	}
}

func calc(board [][]int, x int, y int, width int, height int) int {
	cell := board[y][x]
	liveNum := 0
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= width {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= height {
				continue
			}
			if i == x && j == y {
				continue
			}
			liveNum += board[j][i]
		}
	}
	if cell == 0 {
		if liveNum == 3 {
			return 1
		}
		return 0
	} else {
		if liveNum < 2 {
			return 0
		}
		if liveNum > 3 {
			return 0
		}
		return 1
	}
}
