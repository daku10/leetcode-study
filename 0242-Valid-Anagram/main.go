package main

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}

	return true
}
