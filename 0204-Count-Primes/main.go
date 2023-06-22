package main

import "math"

func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
