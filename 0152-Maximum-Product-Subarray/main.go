package main

import (
	"math"
)

func maxProduct(nums []int) int {
	result := math.MinInt
	for i := 0; i < len(nums); i++ {
		current := 1
		for j := i; j < len(nums); j++ {
			current *= nums[j]
			if result < current {
				result = current
			}
		}
	}
	return result
}
