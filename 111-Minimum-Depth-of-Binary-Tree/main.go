package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := 1
	queue := []*TreeNode{root}
	nextQueue := make([]*TreeNode, 0)
	for {
		if len(queue) == 0 {
			return height
		}
		for _, qn := range queue {
			if qn.Left == nil && qn.Right == nil {
				return height
			}
			if qn.Left != nil {
				nextQueue = append(nextQueue, qn.Left)
			}
			if qn.Right != nil {
				nextQueue = append(nextQueue, qn.Right)
			}
		}
		queue = nextQueue
		nextQueue = make([]*TreeNode, 0)
		height++
	}
}
