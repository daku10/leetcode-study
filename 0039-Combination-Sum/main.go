package main

import (
	"fmt"
	"sort"
)


func combinationSum(candidates []int, target int) [][]int {
	itemMap := make(map[int][][]int)

	for i := 1; i <= target; i++ {
		constructResult(candidates, i, itemMap)
	}

	answerCandidates := itemMap[target]

	if answerCandidates == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	resultMap := make(map[string]bool)

	for i := 0; i < len(answerCandidates); i++ {
		c := answerCandidates[i]
		sort.IntSlice(c).Sort()
		key := fmt.Sprintf("%v", c)
		if !resultMap[key] {
			result = append(result, c)
			resultMap[key] = true
		}
	}

	return result
}

func constructResult(candidates []int, num int, itemMap map[int][][]int) {
	if itemMap[num] != nil {
		return
	}

	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		if num == candidate {
			if itemMap[candidate] == nil {
				itemMap[candidate] = [][]int{{candidate}}
			} else {
				itemMap[candidate] = append(itemMap[candidate], []int{candidate})
			}
		}
	}

	for i := 1; i <= num / 2; i++ {
		n1 := i
		n2 := num - i
		r1 := itemMap[n1]
		r2 := itemMap[n2]
		if r1 == nil || r2 == nil {
			continue
		}
		for j := 0; j < len(r1); j++ {
			for k := 0; k < len(r2); k++ {
				rj := r1[j]
				rk := r2[k]
				tmp := make([]int, 0)
				tmp = append(tmp, rj...)
				tmp = append(tmp, rk...)
				if itemMap[num] == nil {
					im := make([][]int, 0)
					im = append(im, tmp)
					itemMap[num] = im
				} else {				
					itemMap[num] = append(itemMap[num], tmp)
				}
			}
		}
	}
}
