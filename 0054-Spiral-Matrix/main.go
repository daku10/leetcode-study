package main

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	width := len(matrix[0])
	height := len(matrix)
	count := width * height
	x := 0
	y := 0
	xMax := width - 1
	yMax := height - 1
	xMin := 0
	yMin := 0
	order := "x" // x or y or -x or -y

	for i := 0; i < count; i++ {
		result = append(result, matrix[y][x])
		switch order {
		case "x":
			if x == xMax {
				order = "y"
				yMin += 1
			}
		case "y":
			if y == yMax {
				order = "-x"
				xMax -= 1
			}
		case "-x":
			if x == xMin {
				order = "-y"
				yMax -= 1
			}
		case "-y":
			if y == yMin {
				order = "x"
				xMin += 1
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

	return result
}
