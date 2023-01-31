package main

func solve(board [][]byte) {
	width := len(board[0])
	height := len(board)

	// 1000 * y + x
	memo := make(map[int]struct{})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if _, ok := memo[y*1000+x]; ok {
				continue
			}
			if board[y][x] == 'O' {
				// 1000 * y + x
				currentRegion := make(map[int]struct{})
				searchRegion(x, y, board, currentRegion, width, height)
				fillRegion(currentRegion, memo, board, width, height)
			}
		}
	}
}

func searchRegion(x int, y int, board [][]byte, currentRegion map[int]struct{}, width int, height int) {
	if x < 0 {
		return
	}
	if y < 0 {
		return
	}
	if x == width {
		return
	}
	if y == height {
		return
	}
	if board[y][x] == 'X' {
		return
	}
	if _, ok := currentRegion[y*1000+x]; ok {
		return
	}
	currentRegion[y*1000+x] = struct{}{}
	searchRegion(x+1, y, board, currentRegion, width, height)
	searchRegion(x-1, y, board, currentRegion, width, height)
	searchRegion(x, y+1, board, currentRegion, width, height)
	searchRegion(x, y-1, board, currentRegion, width, height)
}

func fillRegion(currentRegion map[int]struct{}, memo map[int]struct{}, board [][]byte, width int, height int) {
	for k := range currentRegion {
		memo[k] = struct{}{}
	}
	for k := range currentRegion {
		y := k / 1000
		if y == 0 || y == height-1 {
			return
		}
		x := k % 1000
		if x == 0 || x == width-1 {
			return
		}
	}
	for k := range currentRegion {
		board[k/1000][k%1000] = 'X'
	}
}
