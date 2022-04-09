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

	return binarySearch(0, max, x)
}

func binarySearch(left int, right int, answer int) int {
	if left == right {
		return left
	}
	if right-left == 1 {
		if right*right == answer {
			return right
		}
		return left
	}
	middle := (left + right) / 2
	sq := middle * middle
	if sq == answer {
		return middle
	} else if sq < answer {
		return binarySearch(middle, right, answer)
	}
	return binarySearch(left, middle, answer)
}
