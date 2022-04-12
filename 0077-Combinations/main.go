package main

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	curr := make([]int, 0)
	backtrack(1, n, k, curr, &res)
	return res
}

func backtrack(start int, n int, k int, curr []int, result *[][]int) {
	if start+k > n+1 {
		return
	}
	if k == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
		return
	}
	for i := start; i <= n; i++ {
		curr = append(curr, i)
		backtrack(i+1, n, k-1, curr, result)
		curr = curr[0 : len(curr)-1]
	}
}
