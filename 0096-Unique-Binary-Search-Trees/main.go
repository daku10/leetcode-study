package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			sum := i - 1
			dp[i] += (dp[sum-j] * dp[j])
		}
	}
	return dp[n]
}
