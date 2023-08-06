package main

import "math"

func majorityElement(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	candidateOne := math.MinInt
	candidateTwo := math.MinInt
	oneNumber := 0
	twoNumber := 0

	for _, num := range nums {
		if candidateOne == num {
			oneNumber++
			continue
		}
		if candidateTwo == num {
			twoNumber++
			continue
		}
		if oneNumber == 0 {
			candidateOne = num
			oneNumber++
			continue
		}
		if twoNumber == 0 {
			candidateTwo = num
			twoNumber++
			continue
		}
		oneNumber--
		twoNumber--
	}

	var result []int
	oneCount := 0
	twoCount := 0
	for _, num := range nums {
		if num == candidateOne {
			oneCount++
		} else if num == candidateTwo {
			twoCount++
		}
	}
	if oneCount > len(nums)/3 {
		result = append(result, candidateOne)
	}
	if twoCount > len(nums)/3 {
		result = append(result, candidateTwo)
	}

	return result
}
