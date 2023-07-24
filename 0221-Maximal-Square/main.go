package main

func maximalSquare(matrix [][]byte) int {
	result := 0
	width := len(matrix[0])
	height := len(matrix)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if matrix[y][x] == '1' {
				res := searchSquare(matrix, x, y)
				if res == width || res == height {
					return res * res
				}
				if res*res > result {
					result = res * res
				}
			}
		}
	}
	return result
}

func searchSquare(matrix [][]byte, x int, y int) int {
	width := len(matrix[0])
	height := len(matrix)
	result := 1
	for {
		for i := x; i <= x+result; i++ {
			for j := y; j <= y+result; j++ {
				if i >= width || j >= height {
					return result
				}
				if matrix[j][i] == '0' {
					return result
				}
			}
		}
		result++
	}
}
