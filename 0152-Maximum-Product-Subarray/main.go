package main

import (
	"math"
)

func maxProduct(nums []int) int {
	result := math.MinInt

	// nums contains even negative numbers, this indicates result
	current := 1
	for i := 0; i < len(nums); i++ {
		current *= nums[i]
		if current > result {
			result = current
		}
		if current == 0 {
			current = 1
		}
	}

	// nums contains odd negative numbers, this indicates result
	current = 1
	for i := len(nums) - 1; i >= 0; i-- {
		current *= nums[i]
		if current > result {
			result = current
		}
		if current == 0 {
			current = 1
		}
	}
	return result
}
