package main

func intersection(nums1 []int, nums2 []int) []int {
	memo1 := make(map[int]struct{})
	memo2 := make(map[int]struct{})
	for _, n := range nums1 {
		memo1[n] = struct{}{}
	}
	for _, n := range nums2 {
		if _, ok := memo1[n]; ok {
			memo2[n] = struct{}{}
		}
	}
	res := make([]int, 0, len(memo2))
	for k := range memo2 {
		res = append(res, k)
	}
	return res
}
