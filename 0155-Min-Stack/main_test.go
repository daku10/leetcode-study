package main

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("testcase", func(t *testing.T) {
		stack := Constructor()
		stack.Push(-2)
		stack.Push(0)
		stack.Push(-3)
		if v := stack.GetMin(); v != -3 {
			t.Fatalf("actual: %v want: %v", v, -3)
		}
		stack.Pop()
		if v := stack.Top(); v != 0 {
			t.Fatalf("actual: %v want: %v", v, 0)
		}
		if v := stack.GetMin(); v != -2 {
			t.Fatalf("actual: %v want: %v", v, -2)
		}
	})
}
