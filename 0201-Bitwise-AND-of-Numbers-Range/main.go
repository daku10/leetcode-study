package main

func rangeBitwiseAnd(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result &= i
		if result == 0 {
			return 0
		}
	}
	return result
}
