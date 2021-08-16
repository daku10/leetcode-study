package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, len(s)) 
	}

	l := len(s)
	for i := 0; i < l; i++ {
		x, y := calcIndex(i, numRows)
		matrix[y][x] = s[i]
	}

	result := make([]byte, 0, l)

	for i := 0; i < numRows; i++ {
		for j := 0; j < l; j++ {
			if matrix[i][j] != 0 {
				result = append(result, matrix[i][j])
			}
		}
	}

	return string(result)
}

func calcIndex(i int, numRows int) (int, int) {
	periodNum := numRows * 2 - 2
	periodWidth := 1 + numRows - 2
	sector := (i / periodNum)
	amari := i % periodNum
	xamari := amari
	if xamari >= numRows {
		xamari = amari - numRows + 1
	} else {
		xamari = 0
	}
	x := sector * periodWidth + xamari

	yamari := amari
	y := 0
	if yamari >= numRows {
		y = (numRows - 1) - (yamari - numRows + 1)
	} else {
		y = yamari
	}
	return x, y
}
