package main

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := counts(i)
		result[i] = count
	}
	return result
}

func counts(n int) int {
	i := 1
	count := 0
	for i <= n {
		if i&n > 0 {
			count++
		}
		i <<= 1
	}
	return count
}
