package main

import "fmt"

func getHint(secret string, guess string) string {
	var bulls, cows int
	memo := make(map[int]struct{})
	secretMap := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
			memo[i] = struct{}{}
		} else {
			secretMap[secret[i]]++
		}
	}

	for i := 0; i < len(secret); i++ {
		if _, ok := memo[i]; ok {
			continue
		}
		if n := secretMap[guess[i]]; n > 0 {
			cows++
			secretMap[guess[i]]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
