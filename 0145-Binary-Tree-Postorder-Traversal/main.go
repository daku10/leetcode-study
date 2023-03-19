package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		result = append(result, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, postorderTraversal(root.Right)...)
	}
	result = append(result, root.Val)
	return result
}
