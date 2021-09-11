package main

import "sort"

func nextPermutation(nums []int)  {

	length := len(nums)

	for i := length - 1; i >= 0; i-- {
		target := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < target {
				nums[i] = nums[j]
				nums[j] = target
				sort.IntSlice(nums[j + 1:length]).Sort()				
				return
			}
		}
	}
	
	// reach this line, it means array is ordered by decending.
	// in that case, next permutation is array ordered by ascending.
	for i := 0; i < length / 2; i++ {
		tmp1 := nums[i]
		nums[i] = nums[length - i - 1]
		nums[length - 1 - i] = tmp1
	}
}
