package main

import (
	"fmt"
	"sort"
)


func combinationSum(candidates []int, target int) [][]int {
	itemMap := make(map[int][][]int)
	resultMap := make(map[string]bool)
	return search(candidates, target, itemMap, resultMap)
}

func search(candidates []int, num int, itemMap map[int][][]int, resultMap map[string]bool) [][]int {
	if itemMap[num] != nil {
		return itemMap[num]
	}

	result := make([][]int, 0)

	length := len(candidates)
	for i := 0; i < length; i++ {
		candidate := candidates[i]
		if candidate == num {
			r := []int{num}
			result = append(result, r)
			if itemMap[num] == nil {
				tmp := make([][]int, 0)
				tmp = append(tmp, r)
				itemMap[num] = tmp
			} else {				
				itemMap[num] = append(itemMap[num], r)
			}		
		}
	}


	for i := 1; i <= num / 2; i++ {
		n1 := i
		r1 := search(candidates, n1, itemMap, resultMap)
		n2 := num - i
		r2 := search(candidates, n2, itemMap, resultMap)
		for j := 0; j < len(r1); j++ {
			for k := 0; k < len(r2); k++ {
				rj := r1[j]
				rk := r2[k]
				tmp := make([]int, 0)
				tmp = append(tmp, rj...)
				tmp = append(tmp, rk...)
				sort.IntSlice(tmp).Sort()
				key := fmt.Sprintf("%d_%v", num, tmp)
				if !resultMap[key] {
					result = append(result, tmp)
					if itemMap[num] == nil {
						im := make([][]int, 0)
						im = append(im, tmp)
						itemMap[num] = im
					} else {				
						itemMap[num] = append(itemMap[num], tmp)
					}
					resultMap[key] = true
				} 
			}
		}
	}

	return result
}
