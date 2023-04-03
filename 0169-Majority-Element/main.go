package main

func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
			count++
		} else {
			if candidate == n {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
