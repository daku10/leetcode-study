package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumSub(root, targetSum, 0)
}

func hasPathSumSub(node *TreeNode, targetSum int, currentSum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return currentSum+node.Val == targetSum
	}
	return hasPathSumSub(node.Left, targetSum, currentSum+node.Val) || hasPathSumSub(node.Right, targetSum, currentSum+node.Val)
}
