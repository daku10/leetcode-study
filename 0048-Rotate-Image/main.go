package main

func rotate(matrix [][]int)  {
	// size of square
    n := len(matrix)

	// rotate outer
	height := 0
	for i := 0; i < n - 1; i++ {		
		pairs := findPair(i, height, n)
		tmp := pairs[3]
		matrix[pairs[3].x][pairs[3].y] = matrix[pairs[2].x][pairs[2].y]
		matrix[pairs[2].x][pairs[2].y] = matrix[pairs[1].x][pairs[1].y]
		matrix[pairs[1].x][pairs[1].y] = matrix[pairs[0].x][pairs[0].y]
		matrix[pairs[0].x][pairs[0].y] = matrix[tmp.x][tmp.y]
	}
}

type Point struct {
	x int
	y int
}

func findPair(x int, y int, n int) []Point {
	return nil
}
