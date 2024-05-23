package main

func arrangeCoins(n int) int {
	currentStair := 1
	currentCoin := 0
	for (currentCoin + currentStair) <= n {
		currentCoin += currentStair
		currentStair++
	}
	return currentStair - 1
}
