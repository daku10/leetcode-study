package main

import (
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	memo := make(map[int]int)
	result := make([]int, 0, n)
	result = append(result, 1)
	for i := 0; i < n-1; i++ {
		result = append(result, calcMin(result, primes, memo))
	}
	return result[n-1]
}

func calcMin(arg []int, primes []int, memo map[int]int) int {
	result := math.MaxInt
	primeResults := make([]int, len(primes))
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		index := memo[p]
		target := p * arg[index]
		primeResults[i] = target
		if target < result {
			result = target
		}
	}
	for i := 0; i < len(primes); i++ {
		if primeResults[i] == result {
			memo[primes[i]]++
		}
	}
	return result
}
