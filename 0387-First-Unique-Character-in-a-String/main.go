package main

func firstUniqChar(s string) int {
	memo := [26]int{}
	for _, a := range s {
		memo[a-'a']++
	}
	for i, a := range s {
		if v := memo[a-'a']; v == 1 {
			return i
		}
	}
	return -1
}
