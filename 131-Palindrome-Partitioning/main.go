package main

import "strings"

func partition(s string) [][]string {
	memo := make(map[string][][]string)
	memo2 := make(map[string]struct{})
	return createPartition(s, memo, memo2)
}

func isEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func createPartition(s string, memo map[string][][]string, memo2 map[string]struct{}) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	if len(s) == 2 {
		if isPalindrome(s) {
			return [][]string{{s}, {s[:1], s[1:]}}
		}
		return [][]string{{s[:1], s[1:]}}
	}
	if v, ok := memo[s]; ok {
		return v
	}
	result := [][]string{}
	if isPalindrome(s) {
		result = append(result, []string{s})
	}
	for i := 1; i < len(s); i++ {
		before := createPartition(s[:i], memo, memo2)
		after := createPartition(s[i:], memo, memo2)
		for _, b := range before {
			for _, a := range after {
				n := make([]string, len(b)+len(a))
				copy(n, b)
				copy(n[len(b):], a)
				k := toKey(n)
				if _, ok := memo2[k]; !ok {
					result = append(result, n)
					memo2[k] = struct{}{}
				}
			}
		}
	}
	memo[s] = result
	return result
}

func toKey(arg []string) string {
	return strings.Join(arg, "_")
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
