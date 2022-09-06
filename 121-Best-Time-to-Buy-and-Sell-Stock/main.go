package main

func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res := prices[j] - prices[i]
			if res > result {
				result = res
			}
		}
	}

	return result
}
