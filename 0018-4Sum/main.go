package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	length := len(nums)
	memo := make(map[string]bool, length)

	candidate := make([]int, 4)

	result := make([][]int, 0)

	// brute force
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				for l := k + 1; l < length; l++ {
					if nums[i] + nums[j] + nums[k] + nums[l] == target {
						candidate[0] = nums[i]
						candidate[1] = nums[j]
						candidate[2] = nums[k]
						candidate[3] = nums[l]
						sort.IntSlice(candidate).Sort()
						key := fmt.Sprintf("%d_%d_%d_%d", candidate[0], candidate[1], candidate[2], candidate[3])
						if !memo[key] {
							memo[key] = true
							result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						}
					}
				}
			}
		}
	}

	return result
}
