package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	mappedNums := make(map[int][]int)

	for i := 0; i < length; i++ {
		if (mappedNums[nums[i]] == nil) {
			mappedNums[nums[i]] = []int{i}
		} else {
			mappedNums[nums[i]] = append(mappedNums[nums[i]], i)
		}
	}

	memoIJ := make(map[int]map[int]bool, length)
	memoI := make(map[int]bool, length)
	resultMemo := make(map[string]bool, length)
	result := [][]int{}
	// brutefource
	for i := 0; i < length; i++ {
		if (memoI[nums[i]]) {
			continue
		} else {
			memoI[nums[i]] = true
		}
		for j := i + 1; j < length; j++ {
			if (memoIJ[nums[i]] == nil) {
				memoIJ[nums[i]] = make(map[int]bool, length)
			}
			if (memoIJ[nums[i]][nums[j]]) {
				continue
			} else {
				memoIJ[nums[i]][nums[j]] = true
			}
			target := 0 - (nums[i] + nums[j])
			targetArray := mappedNums[target]
			tLength := len(targetArray)
			for ti := 0; ti < tLength; ti++ {
				if targetArray[ti] != i && targetArray[ti] != j {
					candidate := []int{nums[i], nums[j], nums[targetArray[ti]]}
					sort.IntSlice(candidate).Sort()
					key := fmt.Sprintf("%d_%d_%d", candidate[0], candidate[1], candidate[2])
					if (!resultMemo[key]) {
						resultMemo[key] = true
						result = append(result, candidate)
					}
					break
				}
			}
		}
	}
	return result
}
