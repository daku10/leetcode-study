package main

func setZeroes(matrix [][]int) {
	rows := make(map[int]struct{})
	columns := make(map[int]struct{})

	height := len(matrix)
	width := len(matrix[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if matrix[y][x] == 0 {
				rows[x] = struct{}{}
				columns[y] = struct{}{}
			}
		}
	}
	for k := range columns {
		for x := 0; x < width; x++ {
			matrix[k][x] = 0
		}
	}
	for k := range rows {
		for y := 0; y < height; y++ {
			matrix[y][k] = 0
		}
	}
}
