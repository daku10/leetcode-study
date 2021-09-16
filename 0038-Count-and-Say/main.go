package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	num := countAndSay(n - 1)
	var targetChar byte = num[0]
	count := 0
	length := len(num)
	var result strings.Builder
	for i := 0; i < length; i++ {
		current := num[i]
		if current != targetChar {
			result.WriteString(fmt.Sprint(count))
			result.WriteByte(targetChar)
			targetChar = current
			count = 1
		} else {			
			count++
		}
	}
	result.WriteString(fmt.Sprint(count))
	result.WriteByte(targetChar)
	return result.String()
}
