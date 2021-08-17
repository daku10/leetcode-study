package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	// -2147483648 <= x <= 2147483647
	if x == -2147483648 {
		return 0
	}

	isNegative := x < 0
	absX := int(math.Abs(float64(x)))

	absXStr := fmt.Sprint(absX)

	reversedStr := reverseStr(absXStr)
	if isBeyond(reversedStr) {
		return 0
	}

	result, _ := strconv.Atoi(reversedStr)

	if isNegative {
		return result * -1
	}
	return result
}

// reserseString
func reverseStr(x string) string {
	result := make([]byte, len(x))
	l := len(x)
	for i := 0; i < l; i++ {
		result[i] = x[l - i - 1]
	} 
	return string(result)
}

func isBeyond(x string) bool {
	standard := "2147483647"
	lenX := len(x)
	lenStandard := len(standard)
	if lenX > lenStandard {
		return true
	} else if lenX < lenStandard {
		return false
	}
	for i := 0; i < lenX; i++ {
		if x[i] > standard[i] {
			return true
		} else if x[i] < standard[i] {
			return false
		}
	}
	// equals to standard
	return false
}

