package main

import "math"

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	notPrimeMap := make(map[int]struct{})
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if _, ok := notPrimeMap[i]; ok {
			continue
		}
		j := 2
		for {
			o := i * j
			if o >= n {
				break
			}
			notPrimeMap[o] = struct{}{}
			j++
		}
	}
	return n - 2 - len(notPrimeMap)
}
