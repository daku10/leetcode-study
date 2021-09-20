package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {

	candidateMap := make(map[int]int)
	length := len(candidates)
	for i := 0; i < length; i++ {
		candidateMap[candidates[i]]++
	}

	fmt.Println(candidateMap)

	return nil
}

