package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	count := n * n
	x := 0
	y := 0
	minX := 0
	maxX := n - 1
	minY := 0
	maxY := n - 1
	order := "x"

	for i := 1; i <= count; i++ {
		matrix[y][x] = i
		switch order {
		case "x":
			if x == maxX {
				order = "y"
				minY++
			}
		case "y":
			if y == maxY {
				order = "-x"
				maxX--
			}
		case "-x":
			if x == minX {
				order = "-y"
				maxY--
			}
		case "-y":
			if y == minY {
				order = "x"
				minX++
			}
		}
		switch order {
		case "x":
			x++
		case "y":
			y++
		case "-x":
			x--
		case "-y":
			y--
		}
	}
	return matrix
}
