package main

func numSquares(n int) int {
	memo := make([]int, n+1)
	return numFindAndCalc(n, memo)
}

func numFindAndCalc(n int, memo []int) int {
	if n < 4 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	// 1 + 1 + 1 + .... is maximal answer
	answer := n
	for i := 1; i*i <= n; i++ {
		square := i * i
		answer = min(answer, 1+numFindAndCalc(n-square, memo))
	}
	memo[n] = answer
	return answer
}
