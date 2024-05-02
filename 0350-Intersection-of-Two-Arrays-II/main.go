package main

func intersect(nums1 []int, nums2 []int) []int {
	memo := make([]int, 1001)
	for _, n := range nums1 {
		memo[n]++
	}
	result := make([]int, 0, max(len(nums1), len(nums2)))
	for _, n := range nums2 {
		if memo[n] > 0 {
			result = append(result, n)
			memo[n]--
		}
	}
	return result
}
