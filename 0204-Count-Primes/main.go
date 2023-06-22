package main

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	notPrimes := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if !notPrimes[i] {
			count++
			j := 3
			for {
				o := i * j
				if o >= n {
					break
				}
				notPrimes[o] = true
				j += 2
			}
		}
	}
	return count
}
