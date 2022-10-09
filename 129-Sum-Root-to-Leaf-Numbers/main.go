package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersSub(root, 0)
}

func sumNumbersSub(node *TreeNode, current int) int {
	new := current*10 + node.Val
	if node.Left != nil && node.Right != nil {
		return sumNumbersSub(node.Left, new) + sumNumbersSub(node.Right, new)
	}
	if node.Left != nil {
		return sumNumbersSub(node.Left, new)
	}
	if node.Right != nil {
		return sumNumbersSub(node.Right, new)
	}
	return new
}
