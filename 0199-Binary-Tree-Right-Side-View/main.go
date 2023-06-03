package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var newQueue []*TreeNode
		for _, q := range queue {
			if q.Left != nil {
				newQueue = append(newQueue, q.Left)
			}
			if q.Right != nil {
				newQueue = append(newQueue, q.Right)
			}
		}
		result = append(result, queue[len(queue)-1].Val)
		queue = newQueue
	}
	return result
}
