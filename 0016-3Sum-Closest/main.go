package main

import "math"

func threeSumClosest(nums []int, target int) int {
	
	length := len(nums)
	
	memoIJ := make(map[int]bool, length)

	result := math.MaxInt32
	diff := math.MaxInt32

	// brute force
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if memoIJ[nums[i] + nums[j]] {
				continue
			} else {
				memoIJ[nums[i] + nums[j]] = true
			}
			for k := j + 1; k < length; k++ {
				current := nums[i] + nums[j] + nums[k]
				absCurrent := absInt(current - target)
				if diff > absCurrent {
					diff = absCurrent
					result = current
				}
				if absCurrent == 0 {
					return current
				}
			}
		}
	}

	return result
}

func absInt(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
