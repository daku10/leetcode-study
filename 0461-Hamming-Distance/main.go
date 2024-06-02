package main

func hammingDistance(x int, y int) int {
	var result int
	for i := 0; i <= 31; i++ {
		c := 1 << i
		a := x & c
		b := y & c
		if a != b {
			result++
		}
	}
	return result
}
