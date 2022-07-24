package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTSub(root, nil, nil)
}

func isValidBSTSub(node *TreeNode, leftMost *int, rightMost *int) bool {
	if node == nil {
		return true
	}
	if node.Left != nil {
		if node.Val <= node.Left.Val {
			return false
		}
		if leftMost != nil && *leftMost >= node.Left.Val {
			return false
		}
		if !isValidBSTSub(node.Left, leftMost, &node.Val) {
			return false
		}
	}
	if node.Right != nil {
		if node.Val >= node.Right.Val {
			return false
		}
		if rightMost != nil && *rightMost <= node.Right.Val {
			return false
		}
		if !isValidBSTSub(node.Right, &node.Val, rightMost) {
			return false
		}
	}
	return true
}
