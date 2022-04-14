package main

func search(nums []int, target int) bool {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left int, right int, target int) bool {
	if right-left <= 1 {
		if nums[right] == target || nums[left] == target {
			return true
		}
		return false
	}
	middle := (left + right) / 2

	leftVal := nums[left]
	rightVal := nums[right]
	middleVal := nums[middle]
	if middleVal == target {
		return true
	}
	if leftVal < target && target <= middleVal {
		return binarySearch(nums, left, middle, target)
	}
	if middleVal < target && target <= rightVal {
		return binarySearch(nums, middle, right, target)
	}
	return binarySearch(nums, left, middle, target) || binarySearch(nums, middle, right, target)
}
