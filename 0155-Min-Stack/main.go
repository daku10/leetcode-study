package main

import (
	"math"
)

type MinStack struct {
	impl    []int
	current int
	min     int
}

func Constructor() MinStack {
	return MinStack{
		impl:    nil,
		current: 0,
		min:     math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.impl = append(this.impl, val)
	this.current++
	if this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	val := this.Top()
	this.current--
	this.impl = this.impl[:len(this.impl)-1]
	if val == this.min {
		m := math.MaxInt
		for _, v := range this.impl {
			if v < m {
				m = v
			}
		}
		this.min = m
	}
}

func (this *MinStack) Top() int {
	return this.impl[this.current-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
