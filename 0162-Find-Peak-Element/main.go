package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := range nums {
		if i == 0 {
			if nums[0] > nums[1] {
				return 0
			}
		} else if i == len(nums)-1 {
			if nums[len(nums)-1] > nums[len(nums)-2] {
				return len(nums) - 1
			}
		} else {
			if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return 0
}
