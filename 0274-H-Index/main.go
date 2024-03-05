package main

import "slices"

func hIndex(citations []int) int {
	slices.Sort(citations)
	for i := 0; i < len(citations); i++ {
		target := citations[i]
		nums := len(citations) - i
		if target == nums {
			return target
		}
		if target > nums {
			return nums
		}
	}
	return 0
}
