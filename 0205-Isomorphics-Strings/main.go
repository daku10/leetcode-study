package main

func isIsomorphic(s string, t string) bool {
	sTot := make([]byte, 128)
	tTos := make([]byte, 128)
	for i := 0; i < len(s); i++ {
		v := sTot[s[i]]
		w := tTos[t[i]]
		if (v == 0 && w != 0) || (v != 0 && w == 0) {
			return false
		}
		if v != 0 {
			if v != t[i] || w != s[i] {
				return false
			}
		} else {
			sTot[s[i]] = t[i]
			tTos[t[i]] = s[i]
		}
	}
	return true
}
