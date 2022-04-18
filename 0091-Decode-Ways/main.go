package main

var memo map[string]int = make(map[string]int)

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

	if _, ok := memo[s]; ok {
		return memo[s]
	}

	if s[0] == '1' || (s[0] == '2' && (s[1]-48) < 7) {
		tmp := numDecodings(s[1:]) + numDecodings(s[2:])
		memo[s] = tmp
		return tmp
	} else {
		tmp := numDecodings(s[1:])
		memo[s] = tmp
		return tmp
	}
}
