package main

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var x, y, z int
	for i := 1; i < n; i++ {
		dp[i] = min(dp[x]*2, dp[y]*3, dp[z]*5)
		if dp[i] == dp[x]*2 {
			x++
		}
		if dp[i] == dp[y]*3 {
			y++
		}
		if dp[i] == dp[z]*5 {
			z++
		}
	}
	return dp[n-1]
}
