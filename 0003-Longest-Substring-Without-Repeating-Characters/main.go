package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
    result := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {			
			if strings.Contains(s[i:j], s[j:j+1]) {
				if j - i > result {
					result = j - i
				}	
				break;
			}
			// this is ugly... when j reaches last index of s, above determining result is evaluated.
			if j == len(s) - 1 && j + 1 - i > result {
				result = j - i + 1
			}
		}
	}
	return result
}
