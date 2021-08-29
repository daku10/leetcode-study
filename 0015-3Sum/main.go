package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	memo := make(map[string]bool)
	result := [][]int{}
	// brutefource
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j == i {
				continue
			}
			for k := j + 1; k < length; k++ {
				if k == j || k == i {
					continue
				}
				if (nums[i] + nums[j] + nums[k]) == 0 {
					candidate := []int{nums[i], nums[j], nums[k]}
					sort.IntSlice(candidate).Sort()
					key := fmt.Sprintf("%d_%d_%d", candidate[0], candidate[1], candidate[2])
					if (!memo[key]) {
						memo[key] = true
						result = append(result, candidate)
					}
				}
			}
		}
	}
	return result
}
