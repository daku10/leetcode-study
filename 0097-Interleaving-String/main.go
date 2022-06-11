package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if isInterleaveSub(s1, s2, s3, true) {
		return true
	}
	return isInterleaveSub(s1, s2, s3, false)
}

var memo map[string]bool = make(map[string]bool)

func isInterleaveSub(s1 string, s2 string, s3 string, isStr1Turn bool) bool {
	if v, ok := memo[fmt.Sprintf("%s_%s_%s_%v", s1, s2, s3, isStr1Turn)]; ok {
		if !v {
			return false
		}
	}
	s1Len := len(s1)
	s2Len := len(s2)
	s3Len := len(s3)
	if s3Len != s1Len+s2Len {
		return false
	}
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
		count := 1
		for i := 1; i < s1Len; i++ {
			if s1[i] == s3[i] {
				count++
			} else {
				break
			}
		}
		if count == s1Len {
			if s2 == s3[s1Len:s3Len] {
				return true
			}
		}
		res := isInterleaveSub(s1[count:s1Len], s2, s3[count:s3Len], false)
		if res {
			return true
		} else {
			memo[fmt.Sprintf("%s_%s_%s_%v", s1[count:s1Len], s2, s3[count:s3Len], isStr1Turn)] = false
		}
		for i := count - 1; i > 0; i-- {
			res := isInterleaveSub(s1[i:s1Len], s2, s3[i:s3Len], false)
			if res {
				return true
			} else {
				memo[fmt.Sprintf("%s_%s_%s_%v", s1[i:s1Len], s2, s3[i:s3Len], isStr1Turn)] = false
			}
			res = isInterleaveSub(s1[i:s1Len], s2, s3[i:s3Len], true)
			if res {
				return true
			} else {
				memo[fmt.Sprintf("%s_%s_%s_%v", s1[i:s1Len], s2, s3[i:s3Len], isStr1Turn)] = false
			}
		}
	} else {
		if s2[0] != s3[0] {
			return false
		}
		if s2Len == 1 {
			return s1 == s3[1:s3Len]
		}
		count := 1
		for i := 1; i < s2Len; i++ {
			if s2[i] == s3[i] {
				count++
			} else {
				break
			}
		}
		if count == s2Len {
			if s1 == s3[s2Len:s3Len] {
				return true
			}
		}
		res := isInterleaveSub(s1, s2[count:s2Len], s3[count:s3Len], true)
		if res {
			return true
		} else {
			memo[fmt.Sprintf("%s_%s_%s_%v", s1, s2[count:s2Len], s3[count:s3Len], isStr1Turn)] = false
		}
		for i := count - 1; i > 0; i-- {
			res := isInterleaveSub(s1, s2[i:s2Len], s3[i:s3Len], true)
			if res {
				return true
			} else {
				memo[fmt.Sprintf("%s_%s_%s_%v", s1, s2[i:s2Len], s3[i:s3Len], isStr1Turn)] = false
			}
			res = isInterleaveSub(s1, s2[i:s2Len], s3[i:s3Len], false)
			if res {
				return true
			} else {
				memo[fmt.Sprintf("%s_%s_%s_%v", s1, s2[i:s2Len], s3[i:s3Len], isStr1Turn)] = false
			}
		}
	}
	memo[fmt.Sprintf("%s_%s_%s_%v", s1, s2, s3, isStr1Turn)] = false
	return false
}
