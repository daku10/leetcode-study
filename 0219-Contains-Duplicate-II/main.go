package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if i-k > 0 {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}

	return false
}
