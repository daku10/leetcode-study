package main

func calculate(s string) int {
	current := 0
	result := 0
	lastNumber := 0
	var operator byte = ' '
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			current = current*10 + int((s[i] - '0'))
		case '+', '-', '*', '/':
			switch operator {
			case '+':
				result += lastNumber
				lastNumber = current
				current = 0
			case '-':
				result += lastNumber
				lastNumber = -1 * current
				current = 0
			case '*':
				lastNumber *= current
				current = 0
			case '/':
				lastNumber /= current
				current = 0
			default:
				lastNumber = current
				current = 0
			}
			operator = s[i]
		default:
		}
	}
	switch operator {
	case '+':
		result += lastNumber
		lastNumber = current
	case '-':
		result += lastNumber
		lastNumber = -1 * current
	case '*':
		lastNumber *= current
	case '/':
		lastNumber /= current
	default:
		lastNumber = current
	}
	result += lastNumber
	return result
}
