package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)

	dfs([]string{}, s, 4, &result)

	return result
}

func dfs(current []string, rest string, index int, result *[]string) {
	if len(rest) > index*3 {
		return
	}
	if index == 0 && len(rest) == 0 {
		*result = append(*result, strings.Join(current, "."))
		return
	} else if index == 0 {
		return
	} else if len(rest) == 0 {
		return
	}
	current = append(current, rest[0:1])
	dfs(current, rest[1:], index-1, result)
	current = current[0 : len(current)-1]
	if len(rest) >= 2 {
		s := rest[0:2]
		if s[0] != '0' {
			current = append(current, rest[0:2])
			dfs(current, rest[2:], index-1, result)
			current = current[0 : len(current)-1]
		}
	}
	if len(rest) >= 3 {
		s := rest[0:3]
		if i, _ := strconv.Atoi(s); i >= 100 && i <= 255 {
			current = append(current, s)
			dfs(current, rest[3:], index-1, result)
			current = current[0 : len(current)-1]
		}
	}
}
