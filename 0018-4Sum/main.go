package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	length := len(nums)
	memo := make(map[string]bool, length)

	itemMap := make(map[int][]int, length)
    
    for i := 0; i < length; i++ {
        if itemMap[nums[i]] == nil {
			itemMap[nums[i]] = make([]int, 0, length)
			itemMap[nums[i]] = append(itemMap[nums[i]], i)
        } else {
            itemMap[nums[i]] = append(itemMap[nums[i]], i)
        }
    }

	memoIJK := make(map[string]bool, length)

	candidate := make([]int, 4)

	result := make([][]int, 0)

	// brute force
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				tIndex := fmt.Sprintf("%d_%d_%d", nums[i], nums[j], nums[k])
				if memoIJK[tIndex] {
					continue
				}
				
				memoIJK[tIndex] = true

				t := nums[i] + nums[j] + nums[k]

				diff := target - t
				diffArray := itemMap[diff]
				if diffArray == nil {
					continue
				}
				diffLength := len(diffArray)
				for l := 0; l < diffLength; l++ {
					diffIndex := diffArray[l]
					if i != diffIndex && j != diffIndex && k != diffIndex {
						candidate[0] = nums[i]
						candidate[1] = nums[j]
						candidate[2] = nums[k]
						candidate[3] = nums[diffIndex]
						sort.IntSlice(candidate).Sort()
						key := fmt.Sprintf("%d_%d_%d_%d", candidate[0], candidate[1], candidate[2], candidate[3])
						if !memo[key] {
							memo[key] = true
							result = append(result, []int{nums[i], nums[j], nums[k], nums[diffIndex]})
						}
						break
					}
				}
			}
		}
	}

	return result
}
