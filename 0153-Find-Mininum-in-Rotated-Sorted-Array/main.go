package main

func findMin(nums []int) int {
	return findMinBinarySearch(nums, 0, len(nums)-1)
}

func findMinBinarySearch(nums []int, start int, end int) int {
	if start == end {
		return nums[start]
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	mid := (start + end) / 2
	if nums[mid] < nums[start] {
		return findMinBinarySearch(nums, start, mid)
	}
	return findMinBinarySearch(nums, mid+1, end)
}
