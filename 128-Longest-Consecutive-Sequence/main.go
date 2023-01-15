package main

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	var maxRes = 1
	var res = 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			continue
		}
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
