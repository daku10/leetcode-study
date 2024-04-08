package main

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		m := n % 4
		if m != 0 {
			return false
		}
		n /= 4
	}
	return true
}
