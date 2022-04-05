package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	result := math.MinInt
	currentResult := math.MinInt
	length := len(nums)
	for i := 0; i < length; i++ {
		tmp := nums[i]
		if currentResult < 0 && tmp > currentResult {
			currentResult = tmp
		} else {
			currentResult += tmp
		}
		if currentResult > result {
			result = currentResult
		}
	}
	return result
}
