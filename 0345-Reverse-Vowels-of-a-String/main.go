package main

func reverseVowels(s string) string {
	ss := []byte(s)
	i := 0
	j := len(s) - 1
loop:
	for i < j {
		for !isVowel(ss[i]) {
			i++
			if i >= len(s) {
				break loop
			}
		}
		for !isVowel(ss[j]) {
			j--
			if j < 0 {
				break loop
			}
		}
		if i > j {
			break
		}
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	return string(ss)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
