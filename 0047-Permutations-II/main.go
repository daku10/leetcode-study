package main

import "fmt"

func permuteUnique(nums []int) [][]int {

	result := make([][]int, 0)

	duplicateMap := make(map[string]bool)

	inner([]int{}, nums, &result, duplicateMap)

	return result
}

func inner(current []int, rest []int, result *[][]int, memo map[string]bool) {
	if len(rest) == 0 {
		key := fmt.Sprintf("%v", current)
		if !memo[key] {
			*result = append(*result, current)
			memo[key] = true
		}
		return
	}

	restLen := len(rest)
	for i := 0; i < restLen; i++ {
		r := rest[i]
		restRest := make([]int, 0)
		for j := 0; j < restLen; j++ {
			if i != j {
				restRest = append(restRest, rest[j])
			}
		}
		newCurrent := make([]int, 0)
		newCurrent = append(newCurrent, current...)
		newCurrent = append(newCurrent, r)
		inner(newCurrent, restRest, result, memo)
	}
}
