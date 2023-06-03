package main

func maxProfit(prices []int) int {
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		cur := prices[i] - minPrice
		if cur > result {
			result = cur
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return result
}
