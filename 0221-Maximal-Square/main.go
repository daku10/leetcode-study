package main

func maximalSquare(matrix [][]byte) int {
	width := len(matrix[0])
	height := len(matrix)
	dp := make([][]int, height)
	for y := 0; y < height; y++ {
		dp[y] = make([]int, width)
		for x := 0; x < width; x++ {
			dp[y][x] = -1
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dp[y][x] == -1 {
				calcDP(matrix, &dp, x, y)
			}
		}
	}

	max := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dp[y][x] > max {
				max = dp[y][x]
			}
		}
	}
	return max * max
}

func calcDP(matrix [][]byte, dp *[][]int, x int, y int) {
	if matrix[y][x] == '0' {
		(*dp)[y][x] = 0
		return
	}
	width := len(matrix[0])
	height := len(matrix)
	if x+1 == width {
		(*dp)[y][x] = 1
		return
	}
	if y+1 == height {
		(*dp)[y][x] = 1
		return
	}
	if (*dp)[y][x+1] == -1 {
		calcDP(matrix, dp, x+1, y)
	}
	if (*dp)[y+1][x] == -1 {
		calcDP(matrix, dp, x, y+1)
	}
	if (*dp)[y+1][x+1] == -1 {
		calcDP(matrix, dp, x+1, y+1)
	}

	(*dp)[y][x] = min(min((*dp)[y][x+1], (*dp)[y+1][x]), (*dp)[y+1][x+1]) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
