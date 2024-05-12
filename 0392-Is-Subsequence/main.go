package main

func isSubsequence(s string, t string) bool {
	si := 0
	for i := 0; i < len(t); i++ {
		if si == len(s) {
			return true
		}
		if t[i] == s[si] {
			si++
		}
	}
	return si == len(s)
}
