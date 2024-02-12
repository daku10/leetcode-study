package main

func singleNumber(nums []int) []int {
	var diff int
	for _, num := range nums {
		diff ^= num
	}
	rightmost := 1
	for diff&rightmost == 0 {
		rightmost <<= 1
	}
	var a, b int
	for _, num := range nums {
		if num&rightmost == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
