package main

func numDecodings(s string) int {
	if s == "" {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	length := len(s)
	dp := make([]int, length+1)
	dp[0] = 1

	for i := 1; i < length+1; i++ {
		if s[i-1] == '0' {
			if i > 1 && (s[i-2] == '1' || s[i-2] == '2') {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 0
			}
		} else if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1]-48 < 7)) {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(s)]
}
