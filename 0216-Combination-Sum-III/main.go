package main

func combinationSum3(k int, n int) [][]int {
	res := createCombination([]int{}, 1, n, k)
	if res == nil {
		return [][]int{}
	}
	return res
}

func calcMax(n int) int {
	maxixmum := 0
	for i := 9; i > 9-n; i-- {
		maxixmum += i
	}
	return maxixmum
}

func createCombination(current []int, start int, target int, remain int) [][]int {
	if start+remain > 10 {
		return nil
	}
	maximum := calcMax(remain)
	if maximum < target {
		return nil
	}
	if maximum == target {
		res := make([]int, len(current))
		copy(res, current)
		for i := 9 - remain + 1; i <= 9; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}
	if remain == 1 {
		if start <= target && target <= 9 {
			res := make([]int, len(current))
			copy(res, current)
			res = append(res, target)
			return [][]int{res}
		}
		return nil
	}
	var result [][]int
	for i := start; i <= 9; i++ {
		cop := make([]int, len(current))
		copy(cop, current)
		cop = append(cop, i)
		res := createCombination(cop, i+1, target-i, remain-1)
		result = append(result, res...)
	}
	return result
}
