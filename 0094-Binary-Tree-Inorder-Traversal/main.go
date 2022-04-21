package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder left->root->right
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	traverse(root, &result)
	return result
}

func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, result)
	*result = append(*result, node.Val)
	traverse(node.Right, result)
}
