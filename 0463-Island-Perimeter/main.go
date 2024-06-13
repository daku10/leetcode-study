package main

func islandPerimeter(grid [][]int) int {
	var result int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				continue
			}
			r := 4
			if x-1 >= 0 && grid[y][x-1] == 1 {
				r--
			}
			if y-1 >= 0 && grid[y-1][x] == 1 {
				r--
			}
			if x+1 < len(grid[0]) && grid[y][x+1] == 1 {
				r--
			}
			if y+1 < len(grid) && grid[y+1][x] == 1 {
				r--
			}
			result += r
		}
	}
	return result
}
