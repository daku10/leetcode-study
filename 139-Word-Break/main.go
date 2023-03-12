package main

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	failed := make(map[string]struct{})
	return memorizedWordBreak(s, wordDict, failed)
}

func memorizedWordBreak(s string, wordDict []string, failed map[string]struct{}) bool {
	if _, ok := failed[s]; ok {
		return false
	}
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			if memorizedWordBreak(s[len(word):], wordDict, failed) {
				return true
			}
		}
	}

	failed[s] = struct{}{}
	return false
}
