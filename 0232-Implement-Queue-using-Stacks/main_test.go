package main

import "testing"

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	if queue.Peek() != 1 {
		t.Errorf("peek should be 1")
	}
	if queue.Pop() != 1 {
		t.Errorf("pop should be 1")
	}
	if queue.Empty() {
		t.Errorf("empty should be false")
	}
}

func TestQueue2(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	if queue.Pop() != 1 {
		t.Errorf("should be 1")
	}
	queue.Push(5)
	if queue.Pop() != 2 {
		t.Errorf("should be 2")
	}
	if queue.Pop() != 3 {
		t.Errorf("should be 3")
	}
	if queue.Pop() != 4 {
		t.Errorf("should be 4")
	}
	if queue.Pop() != 5 {
		t.Errorf("should be 5")
	}
}
