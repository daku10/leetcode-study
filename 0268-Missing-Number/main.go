package main

func missingNumber(nums []int) int {
	n := make([]int, len(nums)+1)
	for _, num := range nums {
		n[num]++
	}
	for i, v := range n {
		if v == 0 {
			return i
		}
	}
	return 0
}
