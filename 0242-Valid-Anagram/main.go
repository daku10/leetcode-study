package main

func isAnagram(s string, t string) bool {
	memo := make(map[byte]int, 26)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		memo[s[i]]++
		memo[t[i]]--
	}

	for _, v := range memo {
		if v != 0 {
			return false
		}
	}

	return true
}
