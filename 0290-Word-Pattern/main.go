package main

import "strings"

func wordPattern(pattern string, s string) bool {
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	for i, si := range ss {
		b := pattern[i]
		if ps, ok := pMap[b]; ok {
			if ps != si {
				return false
			}
			if b != sMap[si] {
				return false
			}
		} else {
			if _, ok := sMap[si]; ok {
				return false
			}
			pMap[b] = si
			sMap[si] = b
		}
	}
	return true
}
