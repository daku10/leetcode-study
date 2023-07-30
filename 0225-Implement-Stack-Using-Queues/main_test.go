package main

import "testing"

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	if stack.Top() != 2 {
		t.Errorf("top should be 2")
	}
	if stack.Pop() != 2 {
		t.Errorf("pop should be 2")
	}
	if stack.Empty() != false {
		t.Errorf("empty should be false")
	}
}
