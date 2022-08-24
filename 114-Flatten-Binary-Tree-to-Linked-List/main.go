package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dummy := &TreeNode{
		Val: root.Val,
	}
	head := dummy
	preorderTraversal(root, &dummy)
	*root = *head.Right
}

func preorderTraversal(node *TreeNode, result **TreeNode) {
	if node == nil {
		return
	}
	(*result).Right = &TreeNode{
		Val: node.Val,
	}
	*result = (*result).Right
	preorderTraversal(node.Left, result)
	preorderTraversal(node.Right, result)
}
