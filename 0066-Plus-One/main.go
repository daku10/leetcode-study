package main

func plusOne(digits []int) []int {
	length := len(digits)
	result := make([]int, length)
	copy(result, digits)
	result[length-1] += 1
	carryOver := false
	for i := length - 1; i >= 0; i-- {
		if carryOver {
			result[i] += 1
			carryOver = false
		}
		if result[i] == 10 {
			result[i] = 0
			carryOver = true
		}
	}
	if carryOver {
		result = append([]int{1}, result...)
	}
	return result
}
