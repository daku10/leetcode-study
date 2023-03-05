package main

func singleNumber(nums []int) int {
	memo := make(map[int]bool)
	for _, n := range nums {
		memo[n] = !memo[n]
	}

	for k, v := range memo {
		if v {
			return k
		}
	}
	return 0
}
