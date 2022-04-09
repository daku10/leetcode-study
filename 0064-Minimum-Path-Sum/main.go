package main

func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])
	dp := make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < width; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < height; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			min := dp[y-1][x]
			if min > dp[y][x-1] {
				min = dp[y][x-1]
			}
			dp[y][x] = grid[y][x] + min
		}
	}
	return dp[height-1][width-1]
}
