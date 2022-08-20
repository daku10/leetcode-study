package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(calcHeight(root.Left, 0), calcHeight(root.Right, 0)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x, y int) int {
	res := x - y
	if res < 0 {
		return -1 * res
	}
	return res
}

func calcHeight(node *TreeNode, height int) int {
	if node == nil {
		return height
	}
	leftHeight := calcHeight(node.Left, height+1)
	rightHeight := calcHeight(node.Right, height+1)
	return max(leftHeight, rightHeight)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
