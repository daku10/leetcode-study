package main

func majorityElement(nums []int) []int {
	memo := make(map[int]int)
	for _, num := range nums {
		memo[num]++
	}

	var result []int
	for k, v := range memo {
		if v > len(nums)/3 {
			result = append(result, k)
		}
	}
	return result
}
