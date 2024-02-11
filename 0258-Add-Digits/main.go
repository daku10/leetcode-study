package main

import "strconv"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	s := strconv.Itoa(num)
	var n int
	for i := 0; i < len(s); i++ {
		n += (int)(s[i] - 48)
	}
	return addDigits(n)
}
