package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= k {
		memo := make(map[int]struct{})
		for i := 0; i < len(nums); i++ {
			if _, ok := memo[nums[i]]; ok {
				return true
			}
			memo[nums[i]] = struct{}{}
		}
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i - k; j <= i+k; j++ {
			if i == j {
				continue
			}
			if j < 0 {
				continue
			}
			if j >= len(nums) {
				break
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
