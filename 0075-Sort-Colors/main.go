package main

func sortColors(nums []int) {
	qsort(nums, 0, len(nums)-1)
}

func qsort(nums []int, left int, right int) {
	if right-left <= 1 {
		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
			return
		}
		return
	}
	i, j := left, right
	pivot := nums[left]
	i++
	for {
		for nums[i] <= pivot && i < right {
			i++
		}
		for nums[j] > pivot && j > left {
			j--
		}
		if i >= j {
			break
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	nums[left], nums[j] = nums[j], nums[left]
	if j-1 > left {
		qsort(nums, left, j-1)
	}
	if j+1 < right {
		qsort(nums, j+1, right)
	}
}
