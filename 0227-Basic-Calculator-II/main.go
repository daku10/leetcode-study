package main

import (
	"strconv"
	"strings"
)

func calculate(s string) int {
	var current strings.Builder
	result := 0
	lastNumber := 0
	var operator byte = ' '
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			current.WriteByte(s[i])
		case '+', '-', '*', '/':
			switch operator {
			case '+':
				result += lastNumber
				lastNumber = toInt(current.String())
				current.Reset()
			case '-':
				result += lastNumber
				lastNumber = -1 * toInt(current.String())
				current.Reset()
			case '*':
				lastNumber *= toInt(current.String())
				current.Reset()
			case '/':
				lastNumber /= toInt(current.String())
				current.Reset()
			default:
				lastNumber = toInt(current.String())
				current.Reset()
			}
			operator = s[i]
		default:
		}
	}
	switch operator {
	case '+':
		result += lastNumber
		lastNumber = toInt(current.String())
		current.Reset()
	case '-':
		result += lastNumber
		lastNumber = -1 * toInt(current.String())
		current.Reset()
	case '*':
		lastNumber *= toInt(current.String())
		current.Reset()
	case '/':
		lastNumber /= toInt(current.String())
		current.Reset()
	default:
		lastNumber = toInt(current.String())
		current.Reset()
	}
	result += lastNumber
	return result
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
