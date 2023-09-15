package main

func productExceptSelf(nums []int) []int {
	current := 1
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		current *= nums[i-1]
		result[i] = current
	}
	current = 1
	for i := len(nums) - 2; i >= 0; i-- {
		current *= nums[i+1]
		result[i] *= current
	}
	return result
}
