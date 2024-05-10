package main

func canConstruct(ransomNote string, magazine string) bool {
	magMemo := make(map[rune]int, 26)
	for _, m := range magazine {
		magMemo[m]++
	}
	for _, r := range ransomNote {
		if v, ok := magMemo[r]; !ok || v <= 0 {
			return false
		}
		magMemo[r]--
	}
	return true
}
