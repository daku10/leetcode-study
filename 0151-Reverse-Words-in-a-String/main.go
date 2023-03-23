package main

import "strings"

func reverseWords(s string) string {
	trimmed := strings.Trim(s, " ")
	splitted := strings.Split(trimmed, " ")
	var res []string
	for i := len(splitted) - 1; i >= 0; i-- {
		if len(splitted[i]) != 0 {
			res = append(res, splitted[i])
		}
	}
	return strings.Join(res, " ")
}
