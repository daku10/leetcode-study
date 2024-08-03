package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	m := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] = i
	}
label:
	for i := 0; i < len(nums1); i++ {
		for j := m[nums1[i]] + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				result[i] = nums2[j]
				continue label
			}
		}
		result[i] = -1
	}
	return result
}
