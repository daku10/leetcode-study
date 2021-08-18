package main

import (
	"strconv"
)

func myAtoi(s string) int {
	trimmed, isPositive := trim(s)
	numOnly := trimNumOnly(trimmed)
	return clamp(numOnly, isPositive)
}

// bool means isPositive
func trim(s string) (string, bool) {

	sLen := len(s)
	trimmed := ""
	for i := 0; i < sLen; i++ {
		if s[i] != ' ' {
			trimmed = s[i:]
			break;
		}
	}
	if len(trimmed) != 0 {
		if trimmed[0] == '-' {
			return trimmed[1:], false
		} else if trimmed[0] == '+' {
			return trimmed[1:], true
		}
	}
	return trimmed, true
}

func trimNumOnly(s string) string {
	arg := s
	argLen := len(arg)
	result := ""
	// trim zero
	j := 0
	for {		
		if j >= argLen {
			arg = "0"
			break;
		}
		if arg[j] != '0' {
			arg = arg[j:]
			break;
		}
		j++
	}

	argLen = len(arg)
	for i := 0; i < argLen; i++ {
		if '0' <= arg[i] && arg[i] <= '9' {
			result = result + string(arg[i])
		} else {
			return string(result)
		}
	}
	return string(result)
}

func clamp(s string, isPositive bool) int {
	// -2147483648 ~ 2147483647
	sign := -1
	if isPositive {
		sign = 1
	}
	standard := ""
	standardNum := 0
	if isPositive {
		standard = "2147483647"
		standardNum = 2147483647
	} else {
		standard = "2147483648"
		standardNum = -2147483648
	}
	standardLen := len(standard)
	sLen := len(s)
	if sLen > standardLen {
		return standardNum
	} else if sLen < standardLen {
		result, _ := strconv.Atoi(s)
		return result * sign
	}
	for i := 0; i < sLen; i++ {
		if s[i] > standard[i] {
			return standardNum
		} else if s[i] < standard[i] {
			result, _ := strconv.Atoi(s)
			return result * sign
		}
	}
	return standardNum
}

