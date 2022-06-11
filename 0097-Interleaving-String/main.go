package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveSub(s1, s2, "", s3)
}

func isInterleaveSub(s1 string, s2 string, result string, s3 string) bool {
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
	if s1[0] == s3[0] {
		ans := isInterleaveSub(s1[1:s1Len], s2, result+string(s1[0]), s3[1:s3Len])
		if ans {
			return true
		}
	}
	if s2[0] == s3[0] {
		ans := isInterleaveSub(s1, s2[1:s2Len], result+string(s2[0]), s3[1:s3Len])
		if ans {
			return true
		}
	}
	return false
}
