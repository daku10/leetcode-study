package main

func sortColors(nums []int) {
	// selection sort
	length := len(nums)
	for i := 0; i < length; i++ {
		minIndex := i
		min := nums[i]
		for j := i + 1; j < length; j++ {
			if nums[j] < min {
				min = nums[j]
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
