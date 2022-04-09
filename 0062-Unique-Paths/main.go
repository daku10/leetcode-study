package main

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	bottomCount := m - 1
	leftCount := n - 1
	small := bottomCount
	if small > leftCount {
		small = leftCount
	}

	sum := bottomCount + leftCount
	result := 1
	for i := 1; i <= small; i++ {
		result *= (sum - i + 1)
		result /= i
	}
	return result
}
