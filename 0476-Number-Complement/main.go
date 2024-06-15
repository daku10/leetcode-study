package main

func findComplement(num int) int {
	var digits []int
	isProcessing := false
	for i := 30; i >= 0; i-- {
		a := 1 << i
		if !isProcessing && a > num {
			continue
		}
		isProcessing = true
		if num >= a {
			num -= a
			digits = append(digits, 1)
		} else {
			digits = append(digits, 0)
		}
	}
	for i := 0; i < len(digits); i++ {
		if digits[i] == 0 {
			digits[i] = 1
		} else {
			digits[i] = 0
		}
	}
	var result int
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 1 {
			result += (1 << (len(digits) - 1 - i))
		}
	}
	return result
}
