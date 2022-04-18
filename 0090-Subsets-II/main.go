package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.IntSlice.Sort(nums)
	result := [][]int{}
	dfs(nums, 0, []int{}, &result)
	return result
}

func dfs(nums []int, start int, current []int, result *[][]int) {
	tmp := make([]int, len(current))
	copy(tmp, current)
	*result = append(*result, tmp)
	length := len(nums)
	for i := start; i < length; i++ {
		if i >= 1 && i > start && nums[i] == nums[i-1] {
			continue
		}
		current = append(current, nums[i])
		dfs(nums, i+1, current, result)
		current = current[0 : len(current)-1]
	}
}
