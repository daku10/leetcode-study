package main

func findPeakElement(nums []int) int {
	return findPeakElementBinarySearch(nums, 0, len(nums)-1)
}

func findPeakElementBinarySearch(nums []int, start int, end int) int {
	if start == end {
		return end
	}
	if end-start == 1 {
		if nums[end] > nums[start] {
			return end
		}
		return start
	}
	mid := (end + start) / 2
	if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
		return mid
	}
	if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
		return findPeakElementBinarySearch(nums, mid, end)
	}
	return findPeakElementBinarySearch(nums, start, mid)
}
