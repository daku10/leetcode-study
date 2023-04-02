package main

import (
	"math"
	"strings"
)

func convertToTitle(columnNumber int) string {
	// 26^6 < max int < 26^7
	var result strings.Builder
	num := columnNumber
	for i := 6; i >= 1; i-- {
		denom := int(math.Pow(26, float64(i)))
		eq := num / denom
		result.WriteByte(byte(eq) + 64)
		num -= eq * denom
	}

	result.WriteByte(byte(num) + 64)

	return reverse(replaceAtmark(result.String()))
}

func replaceAtmark(s string) string {
	var result strings.Builder
	by := []byte(s)
	for i := len(by) - 1; i >= 0; i-- {
		if by[i] == '@' {
			if i == 0 {
				continue
			}
			t := by[i-1]
			if t == '@' {
				continue
			}
			if t == 'A' {
				result.WriteByte('Z')
				by[i-1] = '@'
				continue
			}
			result.WriteByte('Z')
			result.WriteByte(t - 1)
			i--
		} else {
			result.WriteByte(by[i])
		}
	}

	return result.String()
}

func reverse(s string) string {
	i := 0
	j := len(s) - 1
	b := []byte(s)
	for i < j {
		tmp := b[i]
		b[i] = b[j]
		b[j] = tmp
		i++
		j--
	}
	return string(b)
}
