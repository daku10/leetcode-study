package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	impl    []int
	pointer int
}

func Constructor(root *TreeNode) BSTIterator {
	impl := convertToSlice(root)
	return BSTIterator{
		impl:    impl,
		pointer: 0,
	}
}

func convertToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	if root.Left != nil {
		result = append(result, convertToSlice(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, convertToSlice(root.Right)...)
	}
	return result
}

func (this *BSTIterator) Next() int {
	v := this.impl[this.pointer]
	this.pointer++
	return v
}

func (this *BSTIterator) HasNext() bool {
	return this.pointer < len(this.impl)
}
