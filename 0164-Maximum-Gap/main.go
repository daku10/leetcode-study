package main

import "slices"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	slices.Sort(nums)
	var result = 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > result {
			result = nums[i+1] - nums[i]
		}
	}
	return result
}
