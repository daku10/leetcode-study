package main

func longestPalindrome(s string) int {
	memo := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		memo[s[i]]++
	}
	var result int
	var hasOdd bool
	for _, v := range memo {
		if v%2 == 0 {
			result += v
		} else {
			result += v - 1
			hasOdd = true
		}
	}
	if hasOdd {
		return result + 1
	}
	return result
}
