package main

func partition(s string) [][]string {
	var result [][]string
	candidates := createPartition(s)
L:
	for _, c := range candidates {
		for _, r := range result {
			if isEqual(c, r) {
				continue L
			}
		}
		result = append(result, c)
	}

	return result
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
				n := make([]string, len(b)+len(a))
				copy(n, b)
				copy(n[len(b):], a)
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
