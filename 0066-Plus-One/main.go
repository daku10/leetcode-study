package main

func plusOne(digits []int) []int {
	length := len(digits)
	digits[length-1] += 1
	carryOver := false
	for i := length - 1; i >= 0; i-- {
		if carryOver {
			digits[i] += 1
			carryOver = false
		}
		if digits[i] == 10 {
			digits[i] = 0
			carryOver = true
		}
	}
	if carryOver {
		return append([]int{1}, digits...)
	}
	return digits
}
