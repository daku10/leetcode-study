package main

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		mod := n % 3
		if mod != 0 {
			return false
		}
		n /= 3
	}
	return true
}
