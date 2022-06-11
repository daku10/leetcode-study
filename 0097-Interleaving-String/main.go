package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len+s2Len != len(s3) {
		return false
	}
	dp := make([][]bool, s1Len+1)
	for i := 0; i <= s1Len; i++ {
		dp[i] = make([]bool, s2Len+1)
		for j := 0; j <= s2Len; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || dp[i-1][j] && s1[i-1] == s3[i+j-1]
			}
		}
	}

	return dp[s1Len][s2Len]
}
