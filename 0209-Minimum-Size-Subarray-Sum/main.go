package main

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		sum := 0
		var j int
		for j = i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				if (j-i)+1 < result {
					result = (j - i) + 1
				}
				break
			}
		}
		if i == 0 && j == len(nums) {
			return 0
		}
	}
	return result
}
