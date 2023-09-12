package main

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	zeroIndex := -1
	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIndex = i
		}
	}
	if zeroCount >= 2 {
		result := make([]int, len(nums))
		return result
	}
	if zeroCount == 1 {
		result := make([]int, len(nums))
		allProduct := 1
		for i := 0; i < len(nums); i++ {
			if i == zeroIndex {
				continue
			}
			allProduct *= nums[i]
		}
		result[zeroIndex] = allProduct
		return result
	}
	memoPrefix := make([]int, len(nums))
	memoSuffix := make([]int, len(nums))
	current := 1
	for i := len(nums) - 1; i > 0; i-- {
		current *= nums[i]
		memoSuffix[i] = current
	}
	current = 1
	for i := 0; i < len(nums)-1; i++ {
		current *= nums[i]
		memoPrefix[i] = current
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = memoSuffix[i+1]
		} else if i == len(nums)-1 {
			result[i] = memoPrefix[i-1]
		} else {
			result[i] = memoPrefix[i-1] * memoSuffix[i+1]
		}
	}
	return result
}
