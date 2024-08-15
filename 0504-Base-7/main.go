package main

import (
	"strings"
)

func convertToBase7(num int) string {
	maxSeven := 5764801 // 7^9, near 10^7
	if num == 0 {
		return "0"
	}
	var result strings.Builder
	if num < 0 {
		num = -num
		result.WriteByte(45)
	}

	notZeroFlg := false
	for maxSeven > 1 {
		d := num / maxSeven
		num = num % maxSeven
		if d != 0 {
			notZeroFlg = true
		}
		if notZeroFlg {
			result.WriteByte((byte)(d) + 48)
		}
		maxSeven /= 7
	}
	result.WriteByte((byte)(num) + 48)
	return result.String()
}
