package main

import "strings"

func addStrings(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	if len(num1) > len(num2) {
		num2 = strings.Repeat("0", len(num1)-len(num2)) + num2
	} else {
		num1 = strings.Repeat("0", len(num2)-len(num1)) + num1
	}
	return add(num1, num2)
}

func add(num1 string, num2 string) string {
	result := make([]byte, len(num1))
	var carryOver byte = 0
	for i := len(num1) - 1; i >= 0; i-- {
		result[i], carryOver = addByte(num1[i], num2[i], carryOver)
	}
	r := string(result)
	if carryOver == 1 {
		return "1" + r
	}
	return r
}

func addByte(n1 byte, n2 byte, carryOver byte) (byte, byte) {
	r := (n1 - 48) + (n2 - 48) + carryOver
	return (r % 10) + 48, r / 10
}
