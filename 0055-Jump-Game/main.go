package main

func canJump(nums []int) bool {
	result := false
	max := 0
	jump(nums, 0, &result, &max)
	return result
}

func jump(nums []int, index int, result *bool, max *int) {
	if *result {
		return
	}
	current := nums[index]
	if current+index >= len(nums)-1 {
		*result = true
		return
	}
	if current == 0 {
		if *max < index {
			*max = index
		}
		return
	}
	for c := current; c > 0; c-- {
		if index+c+nums[index+c] <= *max {
			continue
		}
		jump(nums, index+c, result, max)
	}
}
