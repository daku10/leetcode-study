package main

func numIslands(grid [][]byte) int {
	width := len(grid[0])
	height := len(grid)
	visited := make([][]byte, height)
	for y := 0; y < height; y++ {
		visited[y] = make([]byte, width)
	}
	result := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if visited[y][x] != 0 {
				continue
			}
			if grid[y][x] == '1' {
				expand(x, y, width, height, grid, visited)
				result++
			}
		}
	}
	return result
}

func expand(x int, y int, width int, height int, grid [][]byte, visited [][]byte) {
	if x < 0 || x >= width {
		return
	}
	if y < 0 || y >= height {
		return
	}
	if visited[y][x] != 0 {
		return
	}
	if grid[y][x] == '0' {
		return
	}
	visited[y][x] = 1
	expand(x-1, y, width, height, grid, visited)
	expand(x+1, y, width, height, grid, visited)
	expand(x, y-1, width, height, grid, visited)
	expand(x, y+1, width, height, grid, visited)
}
