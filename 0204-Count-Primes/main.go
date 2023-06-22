package main

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	notPrimeMap := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if notPrimeMap[i] == 1 {
			continue
		}
		count++
		j := 2
		for {
			o := i * j
			if o >= n {
				break
			}
			notPrimeMap[o] = 1
			j++
		}
	}
	return count
}
