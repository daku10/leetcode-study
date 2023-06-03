package main

func maxProfit(prices []int) int {
	profit := 0
	havingPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > havingPrice {
			profit += prices[i] - havingPrice
			havingPrice = prices[i]
		} else if prices[i] < havingPrice {
			havingPrice = prices[i]
		}
	}
	return profit
}
