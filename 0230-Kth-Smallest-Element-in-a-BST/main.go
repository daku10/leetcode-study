package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	v, _ := find(root, k)
	return v
}

func find(node *TreeNode, count int) (int, int) {
	currentCount := count
	if node.Left != nil {
		v, newCount := find(node.Left, count)
		if newCount == 0 {
			return v, 0
		}
		currentCount = newCount
	}
	currentCount--
	if currentCount == 0 {
		return node.Val, 0
	}
	if node.Right != nil {
		return find(node.Right, currentCount)
	}
	// unreachable
	return 0, currentCount
}
