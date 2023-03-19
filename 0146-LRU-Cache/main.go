package main

import (
	"fmt"
	"math"
)

type LRUCache struct {
	kv map[int]struct {
		val int
		age int
	}
	cap        int
	currentAge int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kv: make(map[int]struct {
			val int
			age int
		}),
		cap:        capacity,
		currentAge: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	this.currentAge++
	if v, ok := this.kv[key]; ok {
		this.kv[key] = struct {
			val int
			age int
		}{
			val: v.val,
			age: this.currentAge,
		}
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.currentAge++
	this.kv[key] = struct {
		val int
		age int
	}{
		val: value,
		age: this.currentAge,
	}
	if len(this.kv) > this.cap {
		min := math.MaxInt
		minK := 0
		for k, v := range this.kv {
			if min > v.age {
				min = v.age
				minK = k
			}
		}
		delete(this.kv, minK)
	}
}

func (this *LRUCache) Debug() {
	fmt.Println("kv")
	fmt.Println(this.kv)
	fmt.Println("currentAge")
	fmt.Println(this.currentAge)
	fmt.Println("cap")
	fmt.Println(this.cap)
}
