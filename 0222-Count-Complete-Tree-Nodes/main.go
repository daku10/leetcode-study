package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		q := queue[0]
		count++
		queue = queue[1:]
		if q.Left != nil {
			queue = append(queue, q.Left)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
		}
	}

	return count
}
