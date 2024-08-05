package main

import "strings"

func findWords(words []string) []string {
	rows := [][]byte{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}
	rowMap := make(map[byte]int)
	for i := 0; i < len(rows); i++ {
		for _, r := range rows[i] {
			rowMap[r] = i
		}
	}

	var result []string
lab:
	for _, word := range words {
		current := -1
		w2 := strings.ToLower(word)
		for i := 0; i < len(w2); i++ {
			r := rowMap[w2[i]]
			if current == -1 {
				current = r
			} else if current != r {
				continue lab
			}
		}
		result = append(result, word)
	}

	return result
}
