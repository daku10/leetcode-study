package main

type MyQueue struct {
	s1 MyStack
	s2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: MyStack{},
		s2: MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}

type MyStack struct {
	m []int
}

func (s *MyStack) Push(x int) {
	s.m = append(s.m, x)
}

func (s *MyStack) Pop() int {
	v := s.m[len(s.m)-1]
	s.m = s.m[:len(s.m)-1]
	return v
}

func (s *MyStack) Peek() int {
	return s.m[len(s.m)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.m) == 0
}
