package main

import (
	"fmt"
	"strings"
)

func convertToBase7(num int) string {
	maxSeven := 5764801 // 7^9, near 10^7
	if num == 0 {
		return "0"
	}
	minusFlg := false
	if num < 0 {
		num = -num
		minusFlg = true
	}

	var result string
	for maxSeven > 1 {
		d := num / maxSeven
		num = num % maxSeven
		result += fmt.Sprint(d)
		maxSeven /= 7
	}
	result += fmt.Sprint(num)
	result = strings.TrimLeft(result, "0")
	if minusFlg {
		result = "-" + result
	}
	return result
}
