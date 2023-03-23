package main

import "math"

func maxProduct(nums []int) int {
	result := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if v := product(nums[i:j]); v > result {
				result = v
			}
		}
	}
	return result
}

func product(nums []int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}
