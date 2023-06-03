package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if i < 2 {
			dp[i] = nums[i]
			continue
		}
		if i == 2 {
			dp[i] = nums[0] + nums[i]
			continue
		}
		dp[i] = nums[i] + max(dp[i-2], dp[i-3])
	}

	return max(dp[len(nums)-1], dp[len(nums)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
