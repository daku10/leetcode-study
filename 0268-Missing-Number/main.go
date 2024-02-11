package main

func missingNumber(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return ((len(nums)+1)*len(nums))/2 - sum
}
