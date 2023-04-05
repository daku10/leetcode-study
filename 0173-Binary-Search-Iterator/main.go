package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodeStack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	node := root
	var stack []*TreeNode
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}
	return BSTIterator{
		nodeStack: stack,
	}
}

func (this *BSTIterator) Next() int {
	node := this.nodeStack[len(this.nodeStack)-1]
	this.nodeStack = this.nodeStack[:len(this.nodeStack)-1]
	n := node.Right
	for n != nil {
		this.nodeStack = append(this.nodeStack, n)
		n = n.Left
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodeStack) > 0
}
