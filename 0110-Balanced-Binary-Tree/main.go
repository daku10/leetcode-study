package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	h := calcHeight(root)
	return h != -1
}

func calcHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := calcHeight(node.Left)
	if left == -1 {
		return -1
	}
	right := calcHeight(node.Right)
	if right == -1 {
		return -1
	}
	diff := abs(left - right)
	if diff > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
