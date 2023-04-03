package main

import "math"

func titleToNumber(columnTitle string) int {
	result := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		b := (columnTitle[i] - 64) // 'A' = 65
		result += int(b) * int(math.Pow(26, float64(len(columnTitle)-1-i)))
	}
	return result
}
