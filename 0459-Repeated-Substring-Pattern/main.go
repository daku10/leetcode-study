package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if repeat(s[:i], s[i:]) {
			return true
		}
	}
	return false
}

func repeat(prefix string, rest string) bool {
	if len(rest)%len(prefix) != 0 {
		return false
	}
	if prefix == rest {
		return true
	}
	if !strings.HasPrefix(rest, prefix) {
		return false
	}
	return repeat(prefix, rest[len(prefix):])
}
