package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	memo := make(map[int]int)
	return maxProfitWithMemo(prices, 0, memo)
}

func maxProfitWithMemo(prices []int, index int, memo map[int]int) int {
	if v, ok := memo[index]; ok {
		return v
	}
	if len(prices) < 2 {
		return 0
	}
	standard := prices[0]
	var candidates []int
	for i := 1; i < len(prices); i++ {
		if prices[i] > standard {
			if i+2 < len(prices) {
				candidates = append(candidates, prices[i]-standard+maxProfitWithMemo(prices[i+2:], index+i+2, memo))
			} else {
				candidates = append(candidates, prices[i]-standard)
			}
		} else {
			candidates = append(candidates, maxProfitWithMemo(prices[i:], index+i, memo))
		}
	}
	if len(candidates) == 0 {
		return 0
	}
	result := candidates[0]
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > result {
			result = candidates[i]
		}
	}
	memo[index] = result
	return result
}
