package main

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	result := 0
	for n > 0 {
		tmp := n / 5
		result += tmp
		n = tmp
	}

	return result
}
