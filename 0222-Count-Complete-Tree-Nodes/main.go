package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := calcHeight(root.Left)
	rightHeight := calcHeight(root.Right)
	if leftHeight == rightHeight {
		// left subtree is full binary tree
		return ((1 << leftHeight) - 1) + 1 + countNodes(root.Right)
	} else {
		// right subtree is full binary tree
		return ((1 << rightHeight) - 1) + 1 + countNodes(root.Left)
	}
}

func calcHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + calcHeight(node.Left)
}
