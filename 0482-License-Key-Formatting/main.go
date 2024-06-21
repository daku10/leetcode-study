package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	ss := strings.ReplaceAll(strings.ToUpper(s), "-", "")
	dashCount := -1
	var result string
	for i := len(ss) - 1; i >= 0; i-- {
		if dashCount == k-1 {
			result = "-" + result
		}
		result = string(ss[i]) + result
		dashCount++
		dashCount %= k
	}
	return result
}
