package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i/2; j++ {
			dp[i] += (dp[i-1-j] * dp[j] * 2)
		}
		if i&1 == 1 {
			dp[i] += dp[(i-1)/2] * dp[(i-1)/2]
		}
	}
	return dp[n]
}
