package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	currentNum := nums[0]
	result := 1
	for i := 1; i < length; i++ {
		if currentNum <= nums[i] {
			for j := i; j < length; j++ {
				if currentNum < nums[j] {
					nums[result] = nums[j]
					currentNum = nums[j]
					i = j
					result++
					break;
				}
			}
		} else {
			currentNum = nums[i]
			result++
		}
	}
	return result
}
