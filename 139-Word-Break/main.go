package main

import "strings"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			if wordBreak(s[len(word):], wordDict) {
				return true
			}
		}
	}

	return false
}
