package main

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums) - 1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	m := (l + r) / 2
	middle := nums[m]
	if middle == target {
		return m
	}
	if l == r {
		if nums[l] > target {
			return l
		}
		return l + 1
	}
	if target < middle {
		return binarySearch(nums, l, m, target)
	}
	return binarySearch(nums, m + 1, r, target)
}
