package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	var sum int
	height := len(matrix)
	width := len(matrix[0])
	sums[0] = make([]int, width)
	for x := 0; x < width; x++ {
		sum += matrix[0][x]
		sums[0][x] = sum
	}
	for y := 1; y < height; y++ {
		sums[y] = make([]int, width)
		var sum int
		for x := 0; x < width; x++ {
			sum += matrix[y][x]
			sums[y][x] = (sum + sums[y-1][x])
		}
	}
	return NumMatrix{matrix: sums}
}

// y, x
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var up, left, upleft int
	if row1 != 0 {
		up = this.matrix[row1-1][col2]
	}
	if col1 != 0 {
		left = this.matrix[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		upleft = this.matrix[row1-1][col1-1]
	}
	return this.matrix[row2][col2] - up - left + upleft
}
