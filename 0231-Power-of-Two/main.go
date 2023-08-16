package main

func isPowerOfTwo(n int) bool {
	for i := 0; i <= 30; i++ {
		if 1<<i == n {
			return true
		}
	}
	return false
}
