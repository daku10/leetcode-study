package main

func reverse(s string) string {
    rs := []rune(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        rs[i], rs[j] = rs[j], rs[i]
    }
    return string(rs)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := s[0:1]
	sLen := len(s)
	for i := 0; i < sLen - 1; i++ {
		for j := 0; j < sLen; j++ {
			prevStart := i - j
			nextEnd := i + j + 2
			prev := ""
			next := ""
			if prevStart >= 0 {
				prev = s[prevStart:i+1]
			} else {
				continue
			}
			if nextEnd < sLen + 1 {
				next = s[i+1:nextEnd]
			}
			if prev == reverse(next) {
				if len(result) < len(prev) * 2 {
					result = prev + next
					continue
				}
			}

			prev = ""
			next = ""
			nextEnd = i + j + 1

			if prevStart >= 0 {
				prev = s[prevStart:i]
			}
			if nextEnd < sLen + 1 {
				next = s[i+1:nextEnd]
			}

			if prev == reverse(next) {
				if len(result) < len(prev) * 2 + 1 {
					result = prev + s[i:i+1] + next
				}
			}
		}
	} 
	return result
}
