package main

func removeDuplicates(nums []int) int {
	length := len(nums)

	// non-existence number in array
	current := -10000
	currentCount := 0
	index := 0
	for i := 0; i < length; i++ {
		if current == nums[i] {
			currentCount++
			if currentCount >= 3 {
				continue
			}
		} else {
			current = nums[i]
			currentCount = 1
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
