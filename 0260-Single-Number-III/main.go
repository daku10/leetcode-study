package main

func singleNumber(nums []int) []int {
	memo := make(map[int]int, len(nums))
	for _, n := range nums {
		memo[n]++
	}
	var res []int
	for k, v := range memo {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
