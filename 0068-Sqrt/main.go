package main

import (
	"math"
)

func mySqrt(x int) int {
	// 2^31 - 1 -> 2^15.5
	max := 46340
	if x == math.MaxInt32 {
		return max
	}
	if x == 0 {
		return 0
	}

	for i := 1; i < max; i++ {
		before := i * i
		if before == x {
			return i
		}
		after := (i + 1) * (i + 1)
		if before < x && x < after {
			return i
		}
	}

	return max
}
