package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[0], max(nums[1], nums[2]))
	}
	dp0 := make([]int, len(nums))
	dp1 := make([]int, len(nums))
	dp0[0] = nums[0]
	dp0[1] = max(nums[0], nums[1])
	dp1[0] = 0
	dp1[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp0[i] = max(dp0[i-1], dp0[i-2]+nums[i])
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}
	return max(dp0[len(nums)-2], dp1[len(nums)-1])
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
