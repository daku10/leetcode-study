package main

import (
	"math"
)

func longestCommonPrefix(strs []string) string {

	minLen := math.MaxInt32
	for _, str := range strs {
		sLen := len(str)
		if minLen > sLen {
			minLen = sLen
		}
	}

	result := ""
	returnFlg := false

	nums := len(strs)
	for i := 0; i <= minLen; i++ {
		prefix := strs[0][:i]
		for j := 1; j < nums; j++ {
			if strs[j][:i] != prefix {
				returnFlg = true
				break
			}
		}
		if (returnFlg) {
			break
		}
		result = prefix
	}

	return result
}
