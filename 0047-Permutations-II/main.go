package main

import "fmt"

func permuteUnique(nums []int) [][]int {

	result := make([][]int, 0)

	inner([]int{}, nums, &result)

	duplicateMap := make(map[string]bool)

	resultLength := len(result)

	removeDuplicate := make([][]int, 0)

	for i := 0; i < resultLength; i++ {
		r := result[i]
		key := fmt.Sprintf("%v", r)
		if !duplicateMap[key] {
			removeDuplicate = append(removeDuplicate, r)
			duplicateMap[key] = true
		}
	}

    return removeDuplicate
}

func inner(current []int, rest []int, result *[][]int) {
	if len(rest) == 0 {
		*result = append(*result, current)
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
		inner(newCurrent, restRest, result)
	}
}
