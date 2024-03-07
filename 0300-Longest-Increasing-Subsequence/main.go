package main

func lengthOfLIS(nums []int) int {
	var result int
	memo := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		res := lengthOfLISInternal(nums[i:], i, memo)
		if res > result {
			result = res
		}
	}
	return result
}

// length, head num
func lengthOfLISInternal(nums []int, index int, memo map[int]int) int {
	if v, ok := memo[index]; ok {
		return v
	}
	if len(nums) == 1 {
		memo[index] = 1
		return 1
	}

	result := 1
	for i := len(nums) - 1; i >= 1; i-- {
		res := lengthOfLISInternal(nums[i:], i+index, memo)
		if nums[0] < nums[i] {
			res++
			if res > result {
				result = res
			}
		}
	}
	memo[index] = result
	return result
}
