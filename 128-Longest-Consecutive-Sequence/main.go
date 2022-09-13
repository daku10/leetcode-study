package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	nums = unique(nums)
	var maxRes = 1
	var res = 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev+1 == nums[i] {
			res++
		} else {
			if maxRes < res {
				maxRes = res
			}
			res = 1
		}
		prev = nums[i]
	}
	if maxRes < res {
		maxRes = res
	}
	return maxRes
}

func unique(nums []int) []int {
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			result = append(result, nums[i])
			prev = nums[i]
		}
	}
	return result
}
