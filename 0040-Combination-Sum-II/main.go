package main

func combinationSum2(candidates []int, target int) [][]int {

	candidateMap := make(map[int]int)
	length := len(candidates)
	for i := 0; i < length; i++ {
		candidateMap[candidates[i]]++
	}

	itemMap := make(map[int][]map[int]int)

	for i := 1; i <= target; i++ {
		constructMap(candidates, i, candidateMap, itemMap)
	}

	result := itemMap[target]
	if result == nil {
		return [][]int{}
	}

	finalResult := make([][]int, 0)

	resultLength := len(result)
	for i := 0; i < resultLength; i++ {
		r := result[i]
		tmp := make([]int, 0)
		for key, value := range r {
			for j := 0; j < value; j++ {
				tmp = append(tmp, key)
			}
		}
		finalResult = append(finalResult, tmp)
	}

	return finalResult
}

func constructMap(candidates []int, target int, candidateMap map[int]int, itemMap map[int][]map[int]int) {
	if (itemMap[target] != nil) {
		return
	}

	result := make([]map[int]int, 0)

	if candidateMap[target] > 0 {
		result = append(result, map[int]int{target: 1})
	}

	// I don't have a good idea.
	duplicateMap := make([]map[int]int, 0)

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
				if success, res := isValid(candidateMap, c1, c2, duplicateMap); success {
					duplicateMap = append(duplicateMap, res)
					result = append(result, res)
				}
			}
		}
	}

	itemMap[target] = result
}

func isValid(candidateMap map[int]int, c1 map[int]int, c2 map[int]int, duplicateMap []map[int]int) (bool, map[int]int) {
	usedMap := make(map[int]int)

	for key, value := range c1 {
		usedMap[key] += value
	}

	for key, value := range c2 {
		usedMap[key] += value
	}

	for key, value := range usedMap {
		if candidateMap[key] < value {
			return false, nil
		}
	}

	lenDuplicate := len(duplicateMap)
	for i := 0; i < lenDuplicate; i++ {
		d := duplicateMap[i]
		isEqual := true
		for key, value := range d {
			if usedMap[key] != value {
				isEqual = false
			}
		}		
		if isEqual {
			return false, nil
		}
	}

	return true, usedMap
}
