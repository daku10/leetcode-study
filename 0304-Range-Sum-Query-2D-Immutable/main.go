package main

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for y := 0; y < len(matrix); y++ {
		sums[y] = make([]int, len(matrix[y]))
		var sum int
		for x := 0; x < len(matrix[y]); x++ {
			sum += matrix[y][x]
			sums[y][x] += sum
		}
	}
	return NumMatrix{matrix: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	if col1 == 0 {
		for y := row1; y <= row2; y++ {
			sum += this.matrix[y][col2]
		}
		return sum
	}
	for y := row1; y <= row2; y++ {
		sum += (this.matrix[y][col2] - this.matrix[y][col1-1])
	}
	return sum
}
