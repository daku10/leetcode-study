package main

func minimumTotal(triangle [][]int) int {
	return dig(triangle, 0, 0, 0)
}

func dig(triangle [][]int, x int, y int, current int) int {
	if len(triangle) == y {
		return current
	}
	current += triangle[y][x]
	return min(dig(triangle, x, y+1, current), dig(triangle, x+1, y+1, current))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
