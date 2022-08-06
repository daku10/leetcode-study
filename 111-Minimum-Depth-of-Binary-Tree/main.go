package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return minDepthSub(root)
}

func minDepthSub(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return min(minDepthSub(root.Left), minDepthSub(root.Right)) + 1
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
