package main

func findDisappearedNumbers(nums []int) []int {
	memo := make([]bool, len(nums))
	for _, n := range nums {
		memo[n-1] = true
	}
	var result []int
	for i, m := range memo {
		if !m {
			result = append(result, i+1)
		}
	}
	return result
}
