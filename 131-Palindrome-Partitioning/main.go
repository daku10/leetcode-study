package main

func partition(s string) [][]string {
	result := [][]string{}
	current := []string{}
	dfs(0, &result, &current, s)
	return result
}

func dfs(start int, result *[][]string, current *[]string, s string) {
	if start >= len(s) {
		res := make([]string, len(*current))
		copy(res, *current)
		*result = append(*result, res)
	}
	for end := start; end < len(s); end++ {
		if isPalindrome(s[start : end+1]) {
			*current = append(*current, s[start:end+1])
			dfs(end+1, result, current, s)
			*current = ((*current)[:len(*current)-1])
		}
	}
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
