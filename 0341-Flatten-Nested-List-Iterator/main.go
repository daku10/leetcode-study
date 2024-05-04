package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type Kind = string

const (
	Int  Kind = "Int"
	List Kind = "List"
)

type NestedInteger struct {
	kind Kind
	i    int
	l    []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.kind == Int
}

func (this NestedInteger) GetInteger() int {
	return this.i
}

func (n *NestedInteger) SetInteger(value int) {
	n.i = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.l = append(this.l, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.l
}

type NestedIterator struct {
	list []int
	p    int
}

func toList(i *NestedInteger) []int {
	if i.IsInteger() {
		return []int{i.GetInteger()}
	}
	var res []int
	for _, l := range i.GetList() {
		res = append(res, toList(l)...)
	}
	return res
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var list []int
	for _, n := range nestedList {
		li := toList(n)
		list = append(list, li...)
	}

	return &NestedIterator{
		list: list,
		p:    0,
	}
}

func (this *NestedIterator) Next() int {
	r := this.list[this.p]
	this.p++
	return r
}

func (this *NestedIterator) HasNext() bool {
	return this.p < len(this.list)
}
