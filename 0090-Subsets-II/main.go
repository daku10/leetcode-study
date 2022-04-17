package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	memo := make(map[string]bool)

	length := len(nums)
	for i := 0; i < length; i++ {
		resLength := len(result)
		for j := 0; j < resLength; j++ {
			tmp := make([]int, len(result[j]))
			copy(tmp, result[j])
			tmp = append(tmp, nums[i])
			sort.IntSlice.Sort(tmp)
			tmpKey := fmt.Sprint(tmp)
			if !memo[tmpKey] {
				result = append(result, tmp)
				memo[tmpKey] = true
			}
		}
	}

	return result
}
