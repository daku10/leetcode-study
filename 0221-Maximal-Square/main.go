package main

func maximalSquare(matrix [][]byte) int {
	width := len(matrix[0])
	height := len(matrix)
	result := 0
	dp := make([][]int, height+1)
	for y := 0; y < height+1; y++ {
		dp[y] = make([]int, width+1)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if matrix[y][x] == '0' {
				continue
			}
			dp[y+1][x+1] = min(min(dp[y][x], dp[y][x+1]), dp[y+1][x]) + 1
			if dp[y+1][x+1] > result {
				result = dp[y+1][x+1]
			}
		}
	}

	return result * result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
