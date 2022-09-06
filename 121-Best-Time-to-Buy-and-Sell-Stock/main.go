package main

import "math"

func maxProfit(prices []int) int {
	result := 0
	cheapestPrice := math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] >= cheapestPrice {
			continue
		}
		cheapestPrice = prices[i]
		for j := i + 1; j < len(prices); j++ {
			res := prices[j] - prices[i]
			if res > result {
				result = res
			}
		}
	}

	return result
}
