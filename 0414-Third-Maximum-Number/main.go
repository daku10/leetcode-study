package main

import (
	"slices"
)

func thirdMax(nums []int) int {
	slices.SortFunc(nums, func(i int, j int) int {
		return j - i
	})
	result := nums[0]
	pos := 1
	for _, n := range nums {
		if n < result {
			result = n
			pos++
			if pos == 3 {
				return result
			}
		}
	}
	return nums[0]
}
