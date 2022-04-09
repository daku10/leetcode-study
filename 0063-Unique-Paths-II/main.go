package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])
	dp := make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	for x := 0; x < width; x++ {
		if obstacleGrid[0][x] == 1 {
			break
		}
		dp[0][x] = 1
	}

	for y := 0; y < height; y++ {
		if obstacleGrid[y][0] == 1 {
			break
		}
		dp[y][0] = 1
	}

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if obstacleGrid[y][x] == 1 {
				dp[y][x] = 0
				continue
			}
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}

	return dp[height-1][width-1]
}
