package main

import (
	"strconv"
	"strings"
)

type stack struct {
	m []int
}

func (s *stack) push(x int) {
	s.m = append(s.m, x)
}

func (s *stack) pop() int {
	r := s.m[len(s.m)-1]
	s.m = s.m[:len(s.m)-1]
	return r
}

func (s *stack) raw() []int {
	return s.m
}

func calculate(s string) int {
	stack := stack{}
	var current strings.Builder
	var operator byte = ' '
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			current.WriteByte(s[i])
		case '+':
			stack.push(toInt(current.String()))
			current.Reset()
			calc(&stack, operator)
			operator = '+'
		case '-':
			stack.push(toInt(current.String()))
			current.Reset()
			calc(&stack, operator)
			operator = '-'
		case '*':
			stack.push(toInt(current.String()))
			current.Reset()
			calc(&stack, operator)
			operator = '*'
		case '/':
			stack.push(toInt(current.String()))
			current.Reset()
			calc(&stack, operator)
			operator = '/'
		default:
		}
	}
	stack.push(toInt(current.String()))
	calc(&stack, operator)
	result := 0
	for _, r := range stack.raw() {
		result += r
	}
	return result
}

func calc(s *stack, operator byte) {
	if operator == ' ' {
		return
	}
	if operator == '+' {
		return
	}
	if operator == '-' {
		r := s.pop()
		s.push(r * -1)
	}
	if operator == '*' {
		r := s.pop()
		t := s.pop()
		s.push(t * r)
	}
	if operator == '/' {
		r := s.pop()
		t := s.pop()
		s.push(t / r)
	}
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
