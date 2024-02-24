package main

import "math"

func numSquares(n int) int {
	memo := make([]int, 10001)
	for i := 1; i <= 100; i++ {
		memo[i*i] = 1
	}
	return numFindAndCalc(n, memo)
}

func numFindAndCalc(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	current := math.MaxInt
	sq := int(math.Sqrt(float64(n)))
	for i := sq; i >= 2; i-- {
		tmp := numFindAndCalc(n-(i*i), memo) + 1
		if tmp == 2 {
			memo[n] = 2
			return 2
		}
		if tmp < current {
			current = tmp
		}
	}
	for i := n / 2; i >= 1; i-- {
		tmp := numFindAndCalc(n-i, memo) + numFindAndCalc(i, memo)
		if tmp == 2 {
			memo[n] = 2
			return 2
		}
		if tmp < current {
			current = tmp
		}
	}
	memo[n] = current
	return current
}
