package main

import (
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	return solve(tokenize(expression))
}

func tokenize(expression string) []string {
	var result []string
	first := 0
	for i, token := range expression {
		if token == '+' || token == '-' || token == '*' {
			result = append(result, expression[first:i])
			result = append(result, string(token))
			first = i + 1
		}
	}
	result = append(result, expression[first:])
	return result
}

func solve(tokens []string) []int {
	if len(tokens) == 1 {
		t, _ := strconv.Atoi(tokens[0])
		return []int{t}
	}
	if len(tokens) == 3 {
		op1, _ := strconv.Atoi(tokens[0])
		op2, _ := strconv.Atoi(tokens[2])
		return []int{calc(op1, op2, tokens[1])}
	}
	var result []int
	for i := 1; i < len(tokens); i += 2 {
		op1s := solve(tokens[:i])
		op2s := solve(tokens[i+1:])
		for _, op1 := range op1s {
			for _, op2 := range op2s {
				result = append(result, calc(op1, op2, tokens[i]))
			}
		}
	}
	return result
}

func calc(op1 int, op2 int, ope string) int {
	switch ope {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	}
	return 0
}
