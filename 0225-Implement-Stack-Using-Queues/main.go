package main

type MyStack struct {
	q1 queue
}

func Constructor() MyStack {
	return MyStack{
		q1: queue{},
	}
}

func (this *MyStack) Push(x int) {
	this.q1.push(x)
}

func (this *MyStack) Pop() int {
	size := this.q1.size()
	for i := 0; i < size-1; i++ {
		this.q1.push(this.q1.pop())
	}
	return this.q1.pop()
}

func (this *MyStack) Top() int {
	size := this.q1.size()
	for i := 0; i < size-1; i++ {
		this.q1.push(this.q1.pop())
	}
	c := this.q1.pop()
	this.q1.push(c)
	return c
}

func (this *MyStack) Empty() bool {
	return this.q1.empty()
}

type queue struct {
	body []int
}

func (q *queue) push(x int) {
	q.body = append(q.body, x)
}

func (q *queue) peek() int {
	return q.body[0]
}

func (q *queue) pop() int {
	c := q.body[0]
	q.body = q.body[1:]
	return c
}

func (q *queue) empty() bool {
	return len(q.body) == 0
}

func (q *queue) size() int {
	return len(q.body)
}
