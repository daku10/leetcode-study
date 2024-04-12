package main

import (
	"math"
)

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	l := math.Log2(float64(n))
	return l == math.Floor(l) && int(l)%2 == 0
}
