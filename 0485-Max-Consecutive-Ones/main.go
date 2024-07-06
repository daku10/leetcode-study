package main

func findMaxConsecutiveOnes(nums []int) int {
	var current, maxCount int
	var isConsecutive bool
	for i := 0; i < len(nums); i++ {
		if isConsecutive {
			if nums[i] == 1 {
				current++
			} else {
				if maxCount < current {
					maxCount = current
				}
				current = 0
				isConsecutive = false
			}
		} else {
			if nums[i] == 1 {
				isConsecutive = true
				current = 1
			}
		}
	}
	if maxCount < current {
		maxCount = current
	}
	return maxCount
}
