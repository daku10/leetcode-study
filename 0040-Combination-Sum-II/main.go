package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {

	candidateMap := make(map[int]int)
	length := len(candidates)
	for i := 0; i < length; i++ {
		candidateMap[candidates[i]]++
	}

	itemMap := make(map[int][][]int)

	for i := 1; i <= target; i++ {
		constructMap(candidates, i, candidateMap, itemMap)
	}

	result := itemMap[target]
	if result == nil {
		return [][]int{}
	}

	duplicatedKeyMap := make(map[string]bool)

	resultLength := len(result)

	finalResult := make([][]int, 0)

	for i := 0; i < resultLength; i++ {
		r := result[i]
		sort.IntSlice(r).Sort()
		key := fmt.Sprintf("%v", r)
		if !duplicatedKeyMap[key] {
			finalResult = append(finalResult, r)
			duplicatedKeyMap[key] = true
		}
	}

	return finalResult
}

func constructMap(candidates []int, target int, candidateMap map[int]int, itemMap map[int][][]int) {
	if (itemMap[target] != nil) {
		return
	}

	result := make([][]int, 0)

	if candidateMap[target] > 0 {
		result = append(result, []int{target})
	}

	for i := 1; i <= target / 2; i++ {
		num1 := i
		num2 := target - i
		r1 := itemMap[num1]
		r2 := itemMap[num2]
		if len(r1) == 0 || len(r2) == 0 {
			continue
		}
		len1 := len(r1)
		len2 := len(r2)
		for j := 0; j < len1; j++ {
			for k := 0; k < len2; k++ {
				c1 := r1[j]
				c2 := r2[k]
				if isValid(candidateMap, c1, c2) {
					tmp := make([]int, 0)
					tmp = append(tmp, c1...)
					tmp = append(tmp, c2...)
					result = append(result, tmp)
				}
			}
		}
	}

	itemMap[target] = result
}

func isValid(candidateMap map[int]int, c1 []int, c2 []int) bool {
	usedMap := make(map[int]int)
	len1 := len(c1)
	for i := 0; i < len1; i++ {
		usedMap[c1[i]]++
	}
	len2 := len(c2)
	for j := 0; j < len2; j++ {
		usedMap[c2[j]]++
	}

	for key, value := range usedMap {
		if candidateMap[key] < value {
			return false
		}
	}

	return true
}
