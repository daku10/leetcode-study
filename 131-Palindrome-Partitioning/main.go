package main

import (
	"reflect"
)

func partition(s string) [][]string {
	var result [][]string
	candidates := createPartition(s)
L:
	for _, c := range candidates {
		for _, r := range result {
			if reflect.DeepEqual(c, r) {
				continue L
			}
		}
		result = append(result, c)
	}

	return result
}

func createPartition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	result := [][]string{}
	if isPalindrome(s) {
		result = append(result, []string{s})
	}
	for i := 1; i < len(s); i++ {
		before := createPartition(s[:i])
		after := createPartition(s[i:])
		for _, b := range before {
			for _, a := range after {
				n := make([]string, 0, len(b)+len(a))
				n = append(n, b...)
				n = append(n, a...)
				result = append(result, n)
			}
		}
	}
	return result
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

// "aab"
// -> {{"aa", "b"}, {"a", "a", "b"}}
// "aa"
// -> {{"aa"}, {"a", "a"}}
// "aaa"
// -> {{"aaa"}, {"aa", "a"}, {"a", "a", "a"}}
// 数学的に組み合わせを作ってpalindromeだけ入れるができそう？

// func createPartition(s string, result *[]string) {
// 	if len(s) == 1 {
// 		return
// 	}
// 	if isPalindrome(s) {
// 		*result = append(*result, s)
// 	}
// 	for i := 1; i < len(s); i++ {
// 		createPartition(s[:i], result)
// 		createPartition(s[i:], result)
// 	}
// }
