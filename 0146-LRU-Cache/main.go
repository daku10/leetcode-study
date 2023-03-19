package main

import (
	"fmt"
)

type Node struct {
	Prev *Node
	Next *Node
	Key  int
}

type LRUCache struct {
	kv   map[int]int
	age  *Node
	kAge map[int]*Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	n := &Node{Key: -1}
	n.Next = n
	n.Prev = n
	return LRUCache{
		kv:   make(map[int]int),
		cap:  capacity,
		age:  n, // dummy
		kAge: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.kv[key]; ok {
		n := this.kAge[key]
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
		n.Next = this.age.Next
		n.Prev = this.age
		this.age.Next.Prev = n
		this.age.Next = n
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.kAge[key]; ok {
		n.Next.Prev = n.Prev
		n.Prev.Next = n.Next
		n.Next = this.age.Next
		n.Prev = this.age
		this.age.Next.Prev = n
		this.age.Next = n
	} else {
		node := &Node{
			Key: key,
		}
		this.kAge[key] = node
		node.Next = this.age.Next
		node.Prev = this.age
		this.age.Next.Prev = node
		this.age.Next = node
	}
	this.kv[key] = value
	if len(this.kv) > this.cap {
		n := this.age.Prev
		k := n.Key
		delete(this.kv, k)
		delete(this.kAge, k)
		n.Prev.Next = this.age
		this.age.Prev = n.Prev
		n.Prev = nil
		n.Next = nil
	}
}

func (this *LRUCache) Debug() {
	fmt.Println("kv")
	fmt.Println(this.kv)
	fmt.Println("cap")
	fmt.Println(this.cap)
	fmt.Println("age next")
	n := this.age
	for {
		fmt.Println(n.Key)
		n = n.Next
		if n == this.age {
			break
		}
	}
	fmt.Println("age prev")
	n = this.age
	for {
		fmt.Println(n.Key)
		n = n.Prev
		if n == this.age {
			break
		}
	}
}
