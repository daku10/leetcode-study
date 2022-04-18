package main

func numDecodings(s string) int {
	if s == "" {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if s[0] == '1' || (s[0] == '2' && (s[1]-48) < 7) {
		return numDecodings(s[1:]) + numDecodings(s[2:])
	} else {
		return numDecodings(s[1:])
	}
}
