package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := Stack{}
	for _, token := range tokens {
		switch token {
		case "+":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b + a)
		case "-":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b - a)
		case "*":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b * a)
		case "/":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b / a)
		default:
			v, _ := strconv.Atoi(token)
			stack.Push(v)
		}
	}
	return stack.Pop()
}

type Stack struct {
	stack []int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}

func (s *Stack) Pop() int {
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}
