package main

import "slices"

func topKFrequent(nums []int, k int) []int {
	memo := make(map[int]int)
	for _, n := range nums {
		memo[n]++
	}
	type s = struct {
		num   int
		count int
	}
	var ss []s
	for k, v := range memo {
		ss = append(ss, struct {
			num   int
			count int
		}{k, v})
	}
	slices.SortFunc(ss, func(a s, b s) int {
		return b.count - a.count
	})
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = ss[i].num
	}
	return result
}
