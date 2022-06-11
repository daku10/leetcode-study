package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if isInterleaveSub(s1, s2, s3, true) {
		return true
	}
	return isInterleaveSub(s1, s2, s3, false)
}

func isInterleaveSub(s1 string, s2 string, s3 string, isStr1Turn bool) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	s3Len := len(s3)
	if s3Len == 0 {
		return s1Len == 0 && s2Len == 0
	}
	if s1Len == 0 {
		return s2 == s3
	}
	if s2Len == 0 {
		return s1 == s3
	}
	if isStr1Turn {
		if s1[0] != s3[0] {
			return false
		}
		if s1Len == 1 {
			return s2 == s3[1:s3Len]
		}
		for i := 1; i < s1Len; i++ {
			if s1[i] != s3[i] {
				return isInterleaveSub(s1[i:s1Len], s2, s3[i:s3Len], false)
			}
			res := isInterleaveSub(s1[i:s1Len], s2, s3[i:s3Len], false)
			if res {
				return true
			}
			res = isInterleaveSub(s1[i:s1Len], s2, s3[i:s3Len], true)
			if res {
				return true
			}
		}
		return s2 == s3[s1Len:s3Len]
	} else {
		if s2[0] != s3[0] {
			return false
		}
		if s2Len == 1 {
			return s1 == s3[1:s3Len]
		}
		for i := 1; i < s2Len; i++ {
			if s2[i] != s3[i] {
				return isInterleaveSub(s1, s2[i:s2Len], s3[i:s3Len], true)
			}
			res := isInterleaveSub(s1, s2[i:s2Len], s3[i:s3Len], true)
			if res {
				return true
			}
			res = isInterleaveSub(s1, s2[i:s2Len], s3[i:s3Len], false)
			if res {
				return true
			}
		}
		return s1 == s3[s2Len:s3Len]
	}
}
