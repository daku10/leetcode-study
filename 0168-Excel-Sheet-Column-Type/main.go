package main

import (
	"strings"
)

func convertToTitle(columnNumber int) string {
	var result strings.Builder
	num := columnNumber
	for num > 0 {
		num--
		rem := num % 26
		result.WriteByte(byte(rem) + 65) // 'A' = 65
		num /= 26
	}
	return reverse(result.String())
}

func reverse(s string) string {
	b := []byte(s)
	i := 0
	j := len(b) - 1
	for i < j {
		tmp := b[i]
		b[i] = b[j]
		b[j] = tmp
		i++
		j--
	}
	return string(b)
}
