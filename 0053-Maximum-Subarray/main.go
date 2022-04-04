package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	result := math.MinInt
	currentResult := 0
	currentMax := math.MinInt
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			currentResult += nums[j]
			if currentMax < currentResult {
				currentMax = currentResult
			}
		}
		currentResult = 0
		if currentMax > result {
			result = currentMax
		}
		currentMax = math.MinInt
	}
	return result
}
