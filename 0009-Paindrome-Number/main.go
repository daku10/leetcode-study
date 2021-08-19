package main

import "fmt"

func isPalindrome(x int) bool {
	xStr := fmt.Sprint(x)
	reversed := reverse(xStr)
	return xStr == reversed
}

func reverse(x string) string {
	xLen := len(x)
	result := make([]byte, xLen)
	for i := 0; i < xLen; i++ {
		result[xLen - i - 1] = x[i]
	}
	return string(result)
}
