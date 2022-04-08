package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	res := strings.Split(strings.Trim(s, " "), " ")
	return len(res[len(res)-1])
}
