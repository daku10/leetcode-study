package main

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	i := -1
	j := 0
	isExpanding := true
	sum := 0
	for j < len(nums) {
		if isExpanding {
			sum += nums[j]
			if sum >= target {
				current := (j - (i + 1)) + 1
				if current == 1 {
					return 1
				}
				if current < result {
					result = current
				}
				isExpanding = false
				i++
			} else {
				j++
			}
		} else {
			sum -= nums[i]
			if sum >= target {
				current := (j - (i + 1)) + 1
				if current == 1 {
					return 1
				}
				if current < result {
					result = current
				}
				i++
			} else {
				isExpanding = true
				j++
			}
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
