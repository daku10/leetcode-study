package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var currentRoot *TreeNode
	currentRoot = root
	for {
		smallest, returnedRoot := findSmallestAndReconstruct(currentRoot)
		k--
		if k == 0 {
			return smallest
		}
		currentRoot = returnedRoot
	}
}

func findSmallestAndReconstruct(root *TreeNode) (int, *TreeNode) {
	node := root
	if node.Left == nil {
		v := node.Val
		newRoot := node.Right
		node.Right = nil
		return v, newRoot
	}
	for node.Left.Left != nil {
		node = node.Left
	}
	v := node.Left.Val
	newLeft := node.Left.Right
	node.Left = newLeft
	return v, root
}
