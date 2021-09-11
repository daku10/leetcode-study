package main

import (
	"sort"
)

func nextPermutation(nums []int)  {

	length := len(nums)

	for i := length - 1; i >= 0; i-- {
		target := nums[i]
		minIndex := -1
		// range of nums[x] is 0 to 100, 101 is higher than maximum abs value.
		// using math.MaxInt is compile error on leetcode environment.
		candidateAbs := 101
		for j := length - 1; j > i; j-- {
			if nums[j] > target && nums[j] - target < candidateAbs  {
				minIndex = j
				candidateAbs = nums[j] - target
			}
		}
		if minIndex != -1 {
			nums[i] = nums[minIndex]
			nums[minIndex] = target
			sort.IntSlice(nums[i + 1:length]).Sort()				
			return
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
