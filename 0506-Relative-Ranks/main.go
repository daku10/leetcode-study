package main

import (
	"fmt"
	"slices"
)

func findRelativeRanks(score []int) []string {
	sorted := make([]int, len(score))
	copy(sorted, score)
	slices.SortFunc(sorted, func(x int, y int) int {
		return y - x
	})
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	memo := make(map[int]string, len(sorted))
	for i := 0; i < len(sorted); i++ {
		if i < 3 {
			memo[sorted[i]] = medals[i]
		} else {
			memo[sorted[i]] = fmt.Sprint(i + 1)
		}
	}
	result := make([]string, 0, len(score))
	for _, s := range score {
		if v, ok := memo[s]; ok {
			result = append(result, v)
		}
	}
	return result
}
