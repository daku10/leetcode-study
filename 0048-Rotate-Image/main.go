package main

func rotate(matrix [][]int)  {
	// size of square
    n := len(matrix)

	for offset := 0; offset < n / 2; offset++ {
		for i := offset; i < n - offset - 1; i++ {		
			pairs := findPair(i, offset, n, offset)
	
			tmp := matrix[pairs[3].y][pairs[3].x]
			matrix[pairs[3].y][pairs[3].x] = matrix[pairs[2].y][pairs[2].x]
			matrix[pairs[2].y][pairs[2].x] = matrix[pairs[1].y][pairs[1].x]
			matrix[pairs[1].y][pairs[1].x] = matrix[pairs[0].y][pairs[0].x]
			matrix[pairs[0].y][pairs[0].x] = tmp
		}
	}
}

type Point struct {
	x int
	y int
}

/*
 return the combination of rotation.
 below is example.
 the same symbol is included in the same conbination.
 x y is point of left corner.
 if x = 0, y = 0, return [{x:0, y:0}, {x:0, y:2}, {x:2, y:2}, {x:0, y:2}]

 o x _ o
 _ $ $ x
 x $ $ _
 o _ x o

 */
func findPair(x int, y int, n int, dan int) []Point {
	result := make([]Point, 0)
	point := Point{x, y}
	result = append(result, point)

	xDiff := x - dan
	
	result = append(result, Point{n - dan - 1, y + xDiff})
	result = append(result, Point{n - dan - 1 - xDiff, n - dan - 1})
	result = append(result, Point{dan, n - dan - 1 - xDiff})
	
	return result
}
