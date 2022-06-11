package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	memo = make([][]int, s1Len+1)
	for i := 0; i <= len(s1); i++ {
		memo[i] = make([]int, s2Len+1)
		for j := 0; j <= s2Len; j++ {
			memo[i][j] = -1
		}
	}
	return isInterleaveSub(s1, 0, s2, 0, "", s3)
}

var memo [][]int

func isInterleaveSub(s1 string, s1Index int, s2 string, s2Index int, result string, s3 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	s3Len := len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}
	if s1Len == 0 {
		return s2 == s3
	}
	if s2Len == 0 {
		return s1 == s3
	}
	if memo[s1Index][s2Index] == 0 {
		return false
	}
	if s1[0] == s3[0] {
		ans := isInterleaveSub(s1[1:s1Len], s1Index+1, s2, s2Index, result+string(s1[0]), s3[1:s3Len])
		if ans {
			return true
		} else {
			memo[s1Index+1][s2Index] = 0
		}
	}
	if s2[0] == s3[0] {
		ans := isInterleaveSub(s1, s1Index, s2[1:s2Len], s2Index+1, result+string(s2[0]), s3[1:s3Len])
		if ans {
			return true
		} else {
			memo[s1Index][s2Index+1] = 0
		}
	}
	memo[s1Index][s2Index] = 0
	return false
}
