package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	lastIndex := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[lastIndex] = nums1[index1]
			index1--
		} else {
			nums1[lastIndex] = nums2[index2]
			index2--
		}
		lastIndex--
	}
	for i := index1; i >= 0; i-- {
		nums1[lastIndex] = nums1[i]
		lastIndex--
	}
	for i := index2; i >= 0; i-- {
		nums1[lastIndex] = nums2[i]
		lastIndex--
	}
}
