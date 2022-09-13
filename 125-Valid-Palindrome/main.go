package main

import "strings"

func isPalindrome(s string) bool {
	sanitaized := sanitaze(s)
	middle := len(sanitaized) / 2
	if len(sanitaized)%2 == 1 {
		middle += 1
	}
	for i, j := 0, len(sanitaized)-1; i < (len(sanitaized) / 2); i, j = i+1, j-1 {
		if sanitaized[i] != sanitaized[j] {
			return false
		}
	}
	return true
}

func sanitaze(s string) string {
	res := strings.ToLower(s)
	var b strings.Builder
	b.Grow(len(res))
	for i := 0; i < len(res); i++ {
		c := res[i]
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			b.WriteByte(c)
		}
	}
	return b.String()
}
