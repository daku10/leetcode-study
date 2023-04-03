package main

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	result := 0
	twoCount := 0
	fiveCount := 0
	for i := 1; i <= n; i++ {
		j := i
		for {
			if j%2 == 0 {
				twoCount++
				j /= 2
			} else {
				break
			}
		}
		j = i
		for {
			if j%5 == 0 {
				fiveCount++
				j /= 5
			} else {
				break
			}
		}
		for twoCount > 0 && fiveCount > 0 {
			result++
			twoCount--
			fiveCount--
		}
	}
	return result
}
