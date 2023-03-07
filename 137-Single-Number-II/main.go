package main

import (
	"math"
)

func singleNumber(nums []int) int {
	current := make([]int, 20)
	negativeCurrent := make([]int, 20)
	negativeCount := 0
	for _, n := range nums {
		var target []int = current
		if n < 0 {
			n = -n
			target = negativeCurrent
			negativeCount++
			negativeCount %= 3
		}
		for i := 19; i >= 0; i-- {
			c := ((int)(math.Pow(3, (float64)(i))))
			if n >= c {
				v := n / c
				target[i] += v
				target[i] %= 3
				n -= (v * c)
			}
		}
	}
	result := 0
	t := current
	if negativeCount != 0 {
		t = negativeCurrent
	}
	for i := 0; i < 20; i++ {
		result += (int)(math.Pow(3, (float64)(i))) * t[i]
	}
	if negativeCount != 0 {
		return result * -1
	}
	return result
}
