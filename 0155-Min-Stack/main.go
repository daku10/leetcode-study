package main

type MinStack struct {
	impl    []int
	min     []int
	current int
}

func Constructor() MinStack {
	return MinStack{
		impl:    nil,
		min:     nil,
		current: 0,
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) Push(val int) {
	this.impl = append(this.impl, val)
	if this.current == 0 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min(this.min[len(this.min)-1], val))
	}
	this.current++
}

func (this *MinStack) Pop() {
	this.current--
	this.impl = this.impl[:len(this.impl)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.impl[this.current-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.current-1]
}
