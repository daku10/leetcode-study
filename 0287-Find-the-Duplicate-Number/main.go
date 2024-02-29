package main

func findDuplicate(nums []int) int {
	memo := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := memo[num]; ok {
			return num
		}
		memo[num] = struct{}{}
	}
	return 0
}
