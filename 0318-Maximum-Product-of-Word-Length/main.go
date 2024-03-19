package main

func maxProduct(words []string) int {
	result := 0
	wordDict := make([][]int, len(words))
	for i := 0; i < len(words); i++ {
		wordDict[i] = make([]int, 26)
		for j := 0; j < len(words[i]); j++ {
			wordDict[i][int(words[i][j]-'a')]++
		}
	}

	for i := 0; i < len(words); i++ {
		currentMax := 0
		for j := i + 1; j < len(words); j++ {
			if currentMax > len(words[j]) {
				continue
			}
			if isNoShare(wordDict, i, j) {
				tmp := len(words[i]) * len(words[j])
				if tmp > result {
					result = tmp
					currentMax = len(words[j])
				}
			}
		}
	}
	return result
}

func isNoShare(wordDict [][]int, i int, j int) bool {
	a := wordDict[i]
	b := wordDict[j]
	for c := 0; c < 26; c++ {
		if a[c] > 0 && b[c] > 0 {
			return false
		}
	}
	return true
}
