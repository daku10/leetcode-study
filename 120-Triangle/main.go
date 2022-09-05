package main

func minimumTotal(triangle [][]int) int {
	for y := len(triangle) - 2; y >= 0; y-- {
		for x := 0; x < len(triangle[y]); x++ {
			triangle[y][x] += min(triangle[y+1][x], triangle[y+1][x+1])
		}
	}

	return triangle[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
