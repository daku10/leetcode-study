package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	result := math.MinInt
	currentResult := 0
	length := len(nums)
	isSearchFirst := true
	for i := 0; i < length; i++ {
		if isSearchFirst {
			currentResult = nums[i]
			isSearchFirst = false
		} else {
			tmp := nums[i]
			if currentResult < 0 && tmp > currentResult {
				currentResult = tmp
			} else if currentResult+tmp < 0 {
				isSearchFirst = true
			} else {
				currentResult += tmp
			}
		}
		if currentResult > result {
			result = currentResult
		}
	}
	return result
}
