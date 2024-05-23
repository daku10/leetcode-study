package main

import "strings"

func countSegments(s string) int {
	ss := strings.Split(s, " ")
	var count int
	for _, sss := range ss {
		if sss != "" {
			count++
		}
	}
	return count
}
