package main

func findKthLargest(nums []int, k int) int {
	if len(nums) < 1000 {
		return bruteforce(nums, k)
	}
	var equalOrLarge []int
	var small []int
	pivot := nums[0]
	for i := 1; i < len(nums); i++ {
		if pivot > nums[i] {
			small = append(small, nums[i])
		} else {
			equalOrLarge = append(equalOrLarge, nums[i])
		}
	}
	if len(equalOrLarge) == k-1 {
		return pivot
	}
	if len(equalOrLarge) >= k {
		return findKthLargest(equalOrLarge, k)
	}
	return findKthLargest(small, k-len(equalOrLarge)-1)
}

func bruteforce(nums []int, k int) int {
	var result int
	for i := 0; i < k; i++ {
		max := nums[0]
		maxIndex := 0
		for j := 0; j < len(nums)-i; j++ {
			if max < nums[j] {
				max = nums[j]
				maxIndex = j
			}
		}
		result = max
		tmp := nums[maxIndex]
		nums[maxIndex] = nums[len(nums)-i-1]
		nums[len(nums)-i-1] = tmp
	}
	return result
}
