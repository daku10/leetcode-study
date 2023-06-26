package main

func isIsomorphic(s string, t string) bool {
	sTot := make(map[byte]byte)
	tTos := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := sTot[s[i]]; ok {
			if v != t[i] {
				return false
			}
			if w, ok := tTos[t[i]]; ok {
				if w != s[i] {
					return false
				}
			}
		} else {
			if _, ok := tTos[t[i]]; ok {
				return false
			}
			sTot[s[i]] = t[i]
			tTos[t[i]] = s[i]
		}
	}
	return true
}
